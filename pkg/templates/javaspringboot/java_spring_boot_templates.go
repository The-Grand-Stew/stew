package javaspringboot

// Add here java code with placeholders

const JavaSpringBootPOMTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<parent>
		<groupId>org.springframework.boot</groupId>
		<artifactId>spring-boot-starter-parent</artifactId>
		<version>2.7.0</version>
		<relativePath/> <!-- lookup parent from repository -->
	</parent>
	<groupId>com.{{ .AppName }}</groupId>
	<artifactId>artifact</artifactId>
	<version>0.0.1-SNAPSHOT</version>
	<name>{{ .AppName }}</name>
	<description>Add service description here</description>
	<properties>
		<java.version>17</java.version>
	</properties>
	<dependencies>
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-web</artifactId>
		</dependency>
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-webflux</artifactId>
		</dependency>

		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-test</artifactId>
			<scope>test</scope>
		</dependency>
		<dependency>
			<groupId>io.projectreactor</groupId>
			<artifactId>reactor-test</artifactId>
			<scope>test</scope>
		</dependency>
	</dependencies>

	<build>
		<plugins>
			<plugin>
				<groupId>org.springframework.boot</groupId>
				<artifactId>spring-boot-maven-plugin</artifactId>
			</plugin>
		</plugins>
	</build>

</project>
`

const JavaSpringBootAppTemplate = `package com.{{ .AppName }};

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class {{ .DomainName }}Application {

	public static void main(String[] args) {
		SpringApplication.run({{ .DomainName }}Application.class, args);
	}

}
`

const JavaSpringBootTestTemplate = `package com.{{ .AppName }};

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class {{ .DomainName }}ApplicationTests {

	@Test
	void contextLoads() {
	}

}
`

const JavaSpringBootControllerTemplate = `package com.{{ .AppName }};

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class {{ .DomainName }}Controller {

	@RequestMapping("/{{ .AppName }}")
	public String get{{ .DomainName }}(){
		return "Hello world!";
	}

	// TODO: add support for other http methods

}
`
