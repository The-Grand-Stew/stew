package javaspringboot

const DatabasePlaceholder = "<!--.DB-->"
const DatabaseDriverProtocolPlaceholder = "{{.DBDriverProtocol}}" // postgresql or mysql

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
		` + DatabasePlaceholder + `
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

const JavaSpringBootApplicationPropertiesBDConnection = `
spring.datasource.url=jdbc:` + DatabaseDriverProtocolPlaceholder + `://${DB_HOST}:${DB_PORT}/${DB_NAME}
spring.datasource.username=${DB_USER}
spring.datasource.password=${DB_PASSWORD}
spring.jpa.show-sql=true
spring.jpa.generate-ddl=true
spring.jpa.hibernate.ddl-auto=update
spring.jpa.properties.hibernate.jdbc.lob.non_contextual_creation=true
`

const JavaSpringBootJpaLibraryTemplate = `
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-data-jpa</artifactId>
		</dependency>
`

const JavaSpringBootPostgresLibraryTemplate = `
		<dependency>
			<groupId>org.postgresql</groupId>
			<artifactId>postgresql</artifactId>
		</dependency>
`

const JavaSpringBootMySqlLibraryTemplate = `
		<dependency>
			<groupId>mysql</groupId>
			<artifactId>mysql-connector-java</artifactId>
		</dependency>
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

const JavaSpringBootDomainTemplate = `package com.{{ .AppName }}.domain;

import javax.persistence.*;
@Entity
@Table(name = "{{ .DomainName | ToLower }}")
public class {{ .DomainName }} {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private long id;
    @Column(name = "name")
    private String name;
    public {{ .DomainName }}() {
    }
    public {{ .DomainName }}(String name) {
        this.name = name;
    }
    public long getId() {
        return id;
    }
    public String getName() {
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }
    @Override
    public String toString() {
        return "{{ .DomainName }} [id=" + id + ", name=" + name + "]";
    }
}
`

const JavaSpringBootRepositoryTemplate = `package com.{{ .AppName }}.repository;

import com.{{ .AppName }}.domain.{{ .DomainName }};
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface {{ .DomainName }}Repository extends JpaRepository<{{ .DomainName }}, Long> {
    List<{{ .DomainName }}> findByName(String name);
}
`

const JavaSpringBootServiceTemplate = `package com.{{ .AppName }}.service;

import com.{{ .AppName }}.domain.{{ .DomainName }};
import com.{{ .AppName }}.repository.{{ .DomainName }}Repository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class {{ .DomainName }}Service {

    {{ .DomainName }}Repository {{ .DomainName | ToLower }}Repository;

    @Autowired
    public {{ .DomainName }}Service({{ .DomainName }}Repository {{ .DomainName | ToLower }}Repository) {
        this.{{ .DomainName | ToLower }}Repository = {{ .DomainName | ToLower }}Repository;
    }

    public void add{{ .DomainName }}({{ .DomainName }} {{ .DomainName | ToLower }}) {
        {{ .DomainName | ToLower }}Repository.save({{ .DomainName | ToLower }});
    }

    public List<{{ .DomainName }}> getAll{{ .DomainName }}s() {
        return {{ .DomainName | ToLower }}Repository.findAll();
    }

    public List<{{ .DomainName }}> getAll{{ .DomainName }}sByName(String name) {
        return {{ .DomainName | ToLower }}Repository.findByName(name);
    }

    public {{ .DomainName }} get{{ .DomainName }}ById(long id) {
        Optional<{{ .DomainName }}> optional{{ .DomainName }}Wrapper = {{ .DomainName | ToLower }}Repository.findById(id);
        return optional{{ .DomainName }}Wrapper.orElse(null);
    }

    public {{ .DomainName }} save({{ .DomainName }} {{ .DomainName | ToLower }}) {
        return {{ .DomainName | ToLower }}Repository.save({{ .DomainName | ToLower }});
    }

    public {{ .DomainName }} update(long id, {{ .DomainName }} {{ .DomainName | ToLower }}) {
        Optional<{{ .DomainName }}> optional{{ .DomainName }}Wrapper = {{ .DomainName | ToLower }}Repository.findById(id);
        if (optional{{ .DomainName }}Wrapper.isPresent()) {
            {{ .DomainName }} _{{ .DomainName | ToLower }} = optional{{ .DomainName }}Wrapper.get();
            _{{ .DomainName | ToLower }}.setName({{ .DomainName | ToLower }}.getName());
            return {{ .DomainName | ToLower }}Repository.save(_{{ .DomainName | ToLower }});
        } else {
            return null;
        }
    }

    public {{ .DomainName }} delete(long id) {
        {{ .DomainName }} {{ .DomainName | ToLower }} = get{{ .DomainName }}ById(id);
        {{ .DomainName | ToLower }}Repository.deleteById(id);
        return {{ .DomainName | ToLower }};
    }

}
`

const JavaSpringBootControllerTemplate = `package com.{{ .AppName }}.controller;

import com.{{ .AppName }}.domain.{{ .DomainName }};
import com.{{ .AppName }}.service.{{ .DomainName }}Service;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
public class {{ .DomainName }}Controller {

	@Autowired
	{{ .DomainName }}Service {{ .DomainName | ToLower }}Service;

	@GetMapping("/{{ .DomainName | ToLower }}")
	public ResponseEntity<List<{{ .DomainName }}>> getAll{{ .DomainName }}s(@RequestParam(required = false) String name) {
		try {
			List<{{ .DomainName }}> {{ .DomainName | ToLower }}s = {{ .DomainName | ToLower }}Service.getAll{{ .DomainName }}s();
			if(name == null)
			{{ .DomainName | ToLower }}s = {{ .DomainName | ToLower }}Service.getAll{{ .DomainName }}s();
			else
			{{ .DomainName | ToLower }}s = {{ .DomainName | ToLower }}Service.getAll{{ .DomainName }}sByName(name);
			return new ResponseEntity<>({{ .DomainName | ToLower }}s, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@GetMapping("/{{ .DomainName | ToLower }}/{id}")
	public ResponseEntity<{{ .DomainName }}> get{{ .DomainName }}ById(@PathVariable("id") long id) {
		try {
			{{ .DomainName }} {{ .DomainName | ToLower }} = {{ .DomainName | ToLower }}Service.get{{ .DomainName }}ById(id);
			return new ResponseEntity<>({{ .DomainName | ToLower }}, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@PostMapping("/{{ .DomainName | ToLower }}")
	public ResponseEntity<{{ .DomainName }}> createTutorial(@RequestBody {{ .DomainName }} {{ .DomainName | ToLower }}) {
		try {
			{{ .DomainName }} _{{ .DomainName | ToLower }} = {{ .DomainName | ToLower }}Service.save({{ .DomainName | ToLower }});
			return new ResponseEntity<>(_{{ .DomainName | ToLower }}, HttpStatus.CREATED);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}

	@PutMapping("/{{ .DomainName | ToLower }}/{id}")
	public ResponseEntity<{{ .DomainName }}> updateTutorial(@PathVariable("id") long id, @RequestBody {{ .DomainName }} {{ .DomainName | ToLower }}) {
		try {
			{{ .DomainName }} _{{ .DomainName | ToLower }} = {{ .DomainName | ToLower }}Service.update(id, {{ .DomainName | ToLower }});
			return new ResponseEntity<>(_{{ .DomainName | ToLower }}, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(null, HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}
	@DeleteMapping("/{{ .DomainName | ToLower }}/{id}")
	public ResponseEntity<{{ .DomainName }}> deleteTutorial(@PathVariable("id") long id) {
		try {
			{{ .DomainName }} _{{ .DomainName | ToLower }} = {{ .DomainName | ToLower }}Service.delete(id);
			return new ResponseEntity<>(_{{ .DomainName | ToLower }}, HttpStatus.OK);
		} catch (Exception e) {
			return new ResponseEntity<>(HttpStatus.INTERNAL_SERVER_ERROR);
		}
	}

}
`
