package fish.eyebrow.stupidsimpleblog.assets

import scala.io.Source

class Asset(resource: String) {
  lazy val content: String = {
    Source.fromResource(s"assets/$resource").mkString
  }
}
