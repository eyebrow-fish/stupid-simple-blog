package fish.eyebrow.stupidsimpleblog.pages

import fish.eyebrow.stupidsimpleblog.assets.Asset
import fish.eyebrow.stupidsimpleblog.entities.Post
import scalatags.Text.TypedTag
import scalatags.Text.all._

import java.time.Instant
import scala.concurrent.{ExecutionContext, Future}

object PostPage extends Page("post", "home") {
  override def assets: Seq[Asset] = Seq(
    new Asset("post.css"),
  )

  override def content(implicit ec: ExecutionContext): Future[TypedTag[String]] = {
    Future.successful {
      val post = Post("How to tie your shoes", Seq("Ask your mother."), "Alexander Johnston", Instant.now())

      div(
        cls := "post",
        div(
          h2(cls := "post-header")(post.title),
          span(post.postedBy),
          span(" at "),
          span(post.postedAt.toString),
        ),
        div(
          cls := "content",
          post.paragraphs.map(p(_)),
        ),
      )
    }
  }
}
