package fish.eyebrow.stupidsimpleblog.api

import akka.actor.ActorSystem
import akka.http.scaladsl.Http
import akka.http.scaladsl.model.StatusCodes
import akka.http.scaladsl.server.Directives._
import com.typesafe.scalalogging.Logger

import scala.concurrent.ExecutionContext

object HttpServer extends App {
  private implicit lazy val logger: Logger = Logger(getClass)
  private implicit lazy val system: ActorSystem = ActorSystem()
  private implicit lazy val ec: ExecutionContext = ExecutionContext.global

  val health = (get & path("health")) {
    complete(StatusCodes.OK)
  }

  Http().newServerAt("localhost", 8080).bind(health).logTime("Start server")
}
