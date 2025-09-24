namespace DywoqGame.Logging.Debugger;

public struct Debugger
{
	/// <summary>
	/// Objects which the debugger will show up in the window.
	/// </summary>
	public List<ITrackable>? Objects { get; set; }

	/// <summary>
	/// The debugger's mode.
	/// </summary>
	public bool On { get; set; }
}
