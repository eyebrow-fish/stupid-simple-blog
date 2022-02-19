package fish.eyebrow.stupidsimpleblog.pages

import akka.http.scaladsl.model.{ContentType, HttpCharsets, HttpEntity, MediaTypes, StatusCodes}
import akka.http.scaladsl.server.Directives.{complete, get}
import akka.http.scaladsl.server.Route
import scalatags.Text.{TypedTag, all}
import scalatags.Text.all.{body, html}

trait Page {
  def render: TypedTag[String]
}

object Page {
  def route(p: Page): Route = {
    get {
      val content = html(all.head(), body(p.render)).render
      val contentType = ContentType(MediaTypes.`text/html`, HttpCharsets.`UTF-8`)
      complete(StatusCodes.OK, HttpEntity(contentType, content))
    }
  }
}
