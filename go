#!/bin/bash -ex
CP=
CP=$CP:/Users/daniel/.m2/repository/org/apache/lucene/lucene-core/5.4.0/lucene-core-5.4.0.jar
CP=$CP:target/luke-with-deps.jar

java -cp $CP org.getopt.luke.Luke 
