package fish.eyebrow.stupidsimpleblog.pages

import scalatags.Text.TypedTag
import scalatags.Text.all._

object Index extends Page {
  override def render: TypedTag[String] = b(123)
}
