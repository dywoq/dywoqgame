namespace DywoqGame.Logging.Stream;

/// <summary>
/// Represents the stream for errors messages.
/// </summary>
/// <param name="maxSize">The max size for buffer.</param>
public class ErrorStream(byte maxSize) : Stream(maxSize)
{

}
