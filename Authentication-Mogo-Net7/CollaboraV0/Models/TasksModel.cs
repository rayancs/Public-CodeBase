using MongoDB.Bson.Serialization.Attributes;

namespace CollaboraV0.Models;

public class TasksModel
{
    [BsonId]
    [BsonRepresentation(MongoDB.Bson.BsonType.ObjectId)]
    public string? TaskId { get; set; }
    [BsonElement("TaskName")]
    public string? TaskName { get; set;}
    public TaskRolesModel? TaskRoles { get; set;}
}
