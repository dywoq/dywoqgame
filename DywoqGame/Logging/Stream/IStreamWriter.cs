namespace DywoqGame.Logging.Stream;

public interface IStreamWriter
{
	/// <summary>
	/// Writes a message to the stream buffer.
	/// If the buffer is overflowed, it removes the last messages to free the space for the message.
	/// </summary>
	/// <param name="message">The message to write.</param>
	public void Write(object message);
}
