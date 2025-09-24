namespace DywoqGame.Logging.Stream;

public class Stream(byte maxSize) : IStreamFlusher, IStreamSize, IStreamWriter, IStreamGet, IDisposable
{
	private List<object> buffer = [];
	private bool disposed = false;

	public object[]? Flush()
	{
		ObjectDisposedException.ThrowIf(disposed, this);
		var flushed = buffer;
		buffer.Clear();
		return flushed?.ToArray();
	}

	public List<object>? GetBuffer()
	{
		ObjectDisposedException.ThrowIf(disposed, this);
		return buffer ?? [];
	}

	public object? GetLastMessage()
	{
		ObjectDisposedException.ThrowIf(disposed, this);
		return buffer.Count == 0 ? null : buffer[^1];
	}

	public byte MaxSize() => maxSize;

	public int Size()
	{
		ObjectDisposedException.ThrowIf(disposed, this);
		return (byte)buffer.Count;
	}

	public void Write(object message)
	{
		ObjectDisposedException.ThrowIf(disposed, this);
		if (buffer.Count >= maxSize)
		{
			buffer.RemoveAt(0);
		}
		buffer.Add(message);
	}

	public override bool Equals(object? obj)
	{
		ObjectDisposedException.ThrowIf(disposed, this);
		return obj is Stream s && Size() == s.Size();
	}

	public override int GetHashCode()
	{
		ObjectDisposedException.ThrowIf(disposed, this);
		return Size().GetHashCode();
	}

	protected virtual void Dispose(bool disposing)
	{
		if (!disposing)
		{
			if (disposed)
			{
				buffer.Clear();
				buffer = null!;
			}
			disposed = true;
		}
	}

	public void Dispose()
	{
		Dispose(true);
		GC.SuppressFinalize(this);
	}

	public static Stream operator +(Stream s, object message)
	{
		s.Write(message);
		return s;
	}

	public static bool operator ==(Stream left, Stream right)
	{
		if (left is null) return right is null;
		return left.Equals(right);
	}

	public static bool operator !=(Stream left, Stream right)
	{
		return !(left == right);
	}
}
