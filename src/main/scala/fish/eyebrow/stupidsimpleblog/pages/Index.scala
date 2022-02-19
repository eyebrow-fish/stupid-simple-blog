package fish.eyebrow.stupidsimpleblog.pages

import fish.eyebrow.stupidsimpleblog.assets.Asset
import scalatags.Text.TypedTag
import scalatags.Text.all._

object Index extends Page {
  override def path: String = ""

  override def content: TypedTag[String] = b("Hello, World!")

  override def assets: Seq[Asset] = Seq(
    new Asset("index.css"),
  )
}
