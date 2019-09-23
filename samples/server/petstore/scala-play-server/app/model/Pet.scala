package model

import play.api.libs.json._

/**
  * A pet for sale in the pet store
  * @param status pet status in the store
  */
@javax.annotation.Generated(value = Array("org.openapitools.codegen.languages.ScalaPlayFrameworkServerCodegen"), date = "2019-09-23T12:38:26.780+02:00[Europe/Prague]")
case class Pet(
  id: Option[Long],
  category: Option[Category],
  name: String,
  photoUrls: List[String],
  tags: Option[List[Tag]],
  status: Option[Pet.Status.Value]
)

object Pet {
  implicit lazy val petJsonFormat: Format[Pet] = Json.format[Pet]

  // noinspection TypeAnnotation
  object Status extends Enumeration {
    val Available = Value("available")
    val Pending = Value("pending")
    val Sold = Value("sold")

    type Status = Value
    implicit lazy val StatusJsonFormat: Format[Value] = Format(Reads.enumNameReads(this), Writes.enumNameWrites[this.type])
  }
}

