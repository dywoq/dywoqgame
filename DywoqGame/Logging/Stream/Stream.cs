namespace DywoqGame.Logging.Stream;

public class Stream(byte maxSize) : IStreamFlusher, IStreamSize, IStreamWriter
{
	public List<object> buffer = [];

	public object[]? Flush()
	{
		var flushed = buffer;
		buffer.Clear();
		return flushed?.ToArray();
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
