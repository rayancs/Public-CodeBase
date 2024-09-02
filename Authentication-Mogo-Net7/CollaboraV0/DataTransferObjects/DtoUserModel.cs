using MongoDB.Bson.Serialization.Attributes;

namespace CollaboraV0.DataTransferObjects;

public class DtoUserModel
{
    public string ?Id { get; set; }
    [BsonElement("EmailId")]
    public string? EmailId { get; set; }

    [BsonElement("Age")]
    public int? Age { get; set; }
}
