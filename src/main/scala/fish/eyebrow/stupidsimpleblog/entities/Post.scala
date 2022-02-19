package fish.eyebrow.stupidsimpleblog.entities

import java.time.Instant

case class Post(title: String, paragraphs: Seq[String], postedBy: String, postedAt: Instant)
