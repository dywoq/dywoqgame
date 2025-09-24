namespace DywoqGame.Logging.Stream;

public interface IStreamGet
{
	/// <summary>
	/// Returns the current stream buffer.
	/// If the stream buffer is null, it returns an empty list.
	/// </summary>
	/// <returns>The current stream buffer.</returns>
	public List<object>? GetBuffer();

	/// <summary>
	/// Returns the latest message from the stream buffer.
	/// If there are no messages, the function returns null.
	/// </summary>
	/// <returns>The latest message.</returns>
	public object? GetLastMessage();
}
