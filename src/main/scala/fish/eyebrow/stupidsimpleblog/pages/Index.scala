package fish.eyebrow.stupidsimpleblog.pages

import fish.eyebrow.stupidsimpleblog.assets.Asset
import fish.eyebrow.stupidsimpleblog.entities.Blog
import scalatags.Text.TypedTag
import scalatags.Text.all._

object Index extends Page("", "home") {
  override def assets: Seq[Asset] = Seq(
    new Asset("index.css"),
    new Asset("blog.css"),
  )

  override def content: TypedTag[String] = {
    val blogs = Seq(Blog("How to tie your shoes", Seq("Ask your mother.")))

    div(
      cls := "blogs",
      blogs.map { blog =>
        div(
          cls := "blog",
          h3(blog.title),
          div(
            cls := "content",
            blog.paragraphs.map(p(_)),
          ),
        )
      }
    )
  }
}
