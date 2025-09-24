namespace DywoqGame.Logging.Stream;

/// <summary>
/// Represents the warning stream for warning messages.
/// </summary>
/// <param name="maxSize">The max size for buffer.</param>
public class WarningStream(byte maxSize) : Stream(maxSize)
{

}
