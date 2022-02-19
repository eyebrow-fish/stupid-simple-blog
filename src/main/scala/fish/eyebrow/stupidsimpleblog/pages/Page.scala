package fish.eyebrow.stupidsimpleblog.pages

import akka.http.scaladsl.model._
import akka.http.scaladsl.server.Directives._
import akka.http.scaladsl.server.Route
import fish.eyebrow.stupidsimpleblog.assets.Asset
import scalatags.Text.all._
import scalatags.Text.{TypedTag, all}

trait Page {
  def path: String

  def content: TypedTag[String]

  def assets: Seq[Asset]

  def render: String = {
    doctype("html")(
      html(
        all.head(
          tag("style")(assets.map(_.content)),
        ),
        body(
          header(
            cls := "main-header",
            a(href := "/")(h3("stupid simple blog")),
          ),
          content,
        ),
      ),
    ).render
  }
}

object Page {
  def route(p: Page): Route = {
    (get & path(p.path)) {
      val contentType = ContentType(MediaTypes.`text/html`, HttpCharsets.`UTF-8`)
      complete(StatusCodes.OK, HttpEntity(contentType, p.render))
    }
  }
}
