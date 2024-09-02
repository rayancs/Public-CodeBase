using MongoDB.Bson.Serialization.Attributes;

namespace CollaboraV0.Models;
public class UserModel
{
    [BsonId]
    [BsonRepresentation(MongoDB.Bson.BsonType.ObjectId)]
    public string? Id { get; set; }
    [BsonElement("EmailId")]
    public string? EmailId { get; set; }

    [BsonElement("Age")]
    public int ? Age { get; set; }
}
