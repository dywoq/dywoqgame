namespace DywoqGame.Logging.Stream;

public class Stream(byte maxSize) : IStreamFlusher, IStreamSize, IStreamWriter, IStreamGet
{
	public List<object> buffer = [];

	public object[]? Flush()
	{
		var flushed = buffer;
		buffer.Clear();
		return flushed?.ToArray();
	}

	public List<object>? GetBuffer()
	{
		return buffer ?? [];
	}

	public object? GetLastMessage()
	{
		return buffer.Count == 0 ? null : buffer[^1];
	}

	public byte MaxSize()
	{
		return maxSize;
	}

	public int Size()
	{
		return (byte)buffer.Count;
	}

	public void Write(object message)
	{
		if (buffer.Count >= maxSize)
		{
			buffer.RemoveAt(0);
		}
		buffer.Add(message);
	}
}
