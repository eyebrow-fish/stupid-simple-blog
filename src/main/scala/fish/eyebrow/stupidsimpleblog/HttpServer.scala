package fish.eyebrow.stupidsimpleblog

import akka.actor.ActorSystem
import akka.http.scaladsl.Http
import akka.http.scaladsl.model._
import akka.http.scaladsl.server.Directives._
import akka.http.scaladsl.server.Route
import com.typesafe.scalalogging.Logger
import fish.eyebrow.stupidsimpleblog.pages._

import scala.concurrent.ExecutionContext

object HttpServer extends App {
  private implicit lazy val logger: Logger = Logger(getClass)
  private implicit lazy val system: ActorSystem = ActorSystem()
  private implicit lazy val ec: ExecutionContext = ExecutionContext.global

  private val route: Route = Page.route(Index) ~
    (get & path("health"))(complete(StatusCodes.OK))

  Http().newServerAt("localhost", 8080).bind(route).logTime("Start server")
}
