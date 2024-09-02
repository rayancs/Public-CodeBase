using CollaboraV0.Models;

namespace CollaboraV0.Interface;

public interface ITaskService
{
    Task<TasksModel> GetTask(string taskId);
    IEnumerable<Task<TasksModel>>GetAllTask();
    Task<bool> DeleteTask(string taskId);
    Task<bool> UpdateTask(string  taskId , TasksModel task);
}
