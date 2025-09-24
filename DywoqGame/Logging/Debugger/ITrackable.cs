namespace DywoqGame.Logging.Debugger;

public interface ITrackable
{
	/// <summary>
	/// Returns the information about the object that implements <see cref="ITrackable"/>.
	/// </summary>
	/// <returns>The information about the object.</returns>
	public ObjectInfo Track();
}
