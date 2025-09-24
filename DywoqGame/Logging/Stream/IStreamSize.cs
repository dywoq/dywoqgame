namespace DywoqGame.Logging.Stream;

public interface IStreamSize
{
	/// <summary>
	/// Returns the current size of the stream buffer.
	/// </summary>
	/// <returns>The stream buffer's size.</returns>
	public int Size();

	/// <summary>
	/// Returns the max size of the stream buffer.
	/// </summary>
	/// <returns>The max size of the stream buffer.</returns>
	public byte MaxSize();
}
