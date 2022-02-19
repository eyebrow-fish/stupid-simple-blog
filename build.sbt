ThisBuild / organization := "fish.eyebrow"
ThisBuild / scalaVersion := "2.13.7"

lazy val root = (project in file("."))
  .settings(
    name := "api",
    libraryDependencies ++= Seq(
      "com.typesafe.akka" %% "akka-http" % "10.2.6",
      "com.typesafe.akka" %% "akka-stream" % "2.6.18",
      "com.typesafe.scala-logging" %% "scala-logging" % "3.9.4",
      "org.slf4j" % "slf4j-simple" % "1.7.36",
      "com.lihaoyi" %% "scalatags" % "0.11.1",
    ),
  )