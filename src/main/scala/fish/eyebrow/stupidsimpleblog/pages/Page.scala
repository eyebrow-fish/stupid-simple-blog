package fish.eyebrow.stupidsimpleblog.pages

import akka.http.scaladsl.model._
import akka.http.scaladsl.server.Directives.{path, _}
import akka.http.scaladsl.server.Route
import fish.eyebrow.stupidsimpleblog.assets.Asset
import scalatags.Text.all._
import scalatags.Text.{TypedTag, all}

import scala.concurrent.{ExecutionContext, Future}

abstract class Page(val uri: String, title: String) {
  def assets: Seq[Asset]

  def content(implicit ec: ExecutionContext): Future[TypedTag[String]]

  def defaultAssets: Seq[Asset] = Seq(new Asset("index.css"))

  private def render(implicit ec: ExecutionContext): Future[String] = {
    content.map { c =>
      doctype("html")(
        html(
          all.head(
            tag("title")(s"stupid simple blog - $title"),
            (defaultAssets ++ assets).map(_.content),
          ),
          body(
            header(
              cls := "main-header",
              a(cls := "header-text", href := "/")(h2("stupid simple blog")),
            ),
            c,
          ),
        ),
      ).render
    }
  }

  def route(implicit ec: ExecutionContext): Route = {
    (get & path(uri)) {
      onSuccess(render) { x =>
        val contentType = ContentType(MediaTypes.`text/html`, HttpCharsets.`UTF-8`)
        complete(StatusCodes.OK, HttpEntity(contentType, x))
      }
    }
  }
}

object Page {
  def all: Seq[Page] = Seq(IndexPage, PostPage)
}
