namespace DywoqGame.Logging.Stream;

public interface IStreamFlusher
{
	/// <summary>
	/// Flushes and returns the stream buffer, if there are messages on the buffer.
	/// </summary>
	/// <returns>The flushes buffer.</returns>
	public object[]? Flush();
}
