namespace DywoqGame.Logging.Debugger;

/// <summary>
/// Represents object info for debugger.
/// </summary>
/// <param name="Name">The name of object.</param>
/// <param name="Fields">The fields for object.</param>
public record ObjectInfo(string Name, Dictionary<string, object> Fields);
