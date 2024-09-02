using DataTransferObjects;
using MongoDB.Bson.Serialization.Attributes;

namespace CollaboraV0.Models;

public class MyPeopleModel
{
    [BsonId]
    [BsonRepresentation(MongoDB.Bson.BsonType.ObjectId)]
    public string? Id { get; set; }
    [BsonElement("Email")]
    public string? Email { get; set; } = null;
    [BsonElement] public List<RequestUserModel>?Friends { get; set;}
}
