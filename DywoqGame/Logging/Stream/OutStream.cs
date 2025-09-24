namespace DywoqGame.Logging.Stream;

/// <summary>
/// Represents the stream for ordinary messages.
/// </summary>
/// <param name="maxSize">The max size for buffer.</param>
public class OutStream(byte maxSize) : Stream(maxSize)
{

}
