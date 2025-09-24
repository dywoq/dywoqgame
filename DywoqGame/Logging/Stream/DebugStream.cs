namespace DywoqGame.Logging.Stream;

/// <summary>
/// Represents the debugging stream for debug messages.
/// </summary>
/// <param name="maxSize">The max size for buffer.</param>
public class DebugStream(byte maxSize) : Stream(maxSize)
{

}
