FROM openjdk:8-jdk-alpine
COPY --from=build /app/build/libs/jpashop-0.0.1-SNAPSHOT.jar
ENTRYPOINT ["java","-Djava.security.egd=file:/dev/./urandom","-jar","/myboot.jar"]
