namespace DywoqGame.Logging.Stream;

/// <summary>
/// Represents the stream for fatal errors messages.
/// </summary>
/// <param name="maxSize">The max size for buffer.</param>
public class FatalStream(byte maxSize) : Stream(maxSize)
{

}
