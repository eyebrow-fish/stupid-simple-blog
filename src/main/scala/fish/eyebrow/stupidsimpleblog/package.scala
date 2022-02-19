package fish.eyebrow

import com.typesafe.scalalogging.Logger

import java.time.Instant
import scala.concurrent.{ExecutionContext, Future}

package object stupidsimpleblog {
  implicit class FutureEx[T](f: Future[T])(implicit ec: ExecutionContext, logger: Logger) {
    def logTime(msg: String): Future[T] = {
      for {
        s <- Future.successful {
          logger.info(s"""Executing "$msg"...""")
          Instant.now.toEpochMilli
        }
        t <- f
      } yield {
        val n = Instant.now.toEpochMilli
        logger.info(s"""Finished "$msg" in ${n - s}ms.""")
        t
      }
    }
  }
}
