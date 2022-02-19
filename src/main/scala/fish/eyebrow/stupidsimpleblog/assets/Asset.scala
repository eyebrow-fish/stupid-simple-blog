package fish.eyebrow.stupidsimpleblog.assets

import scalatags.Text.TypedTag
import scalatags.Text.all._

import scala.io.Source

class Asset(resource: String) {
  lazy val content: TypedTag[String] = {
    tag("style") {
      Source.fromResource(s"assets/$resource").mkString
    }
  }
}
