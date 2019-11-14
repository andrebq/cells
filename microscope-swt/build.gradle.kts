plugins {
    java
    application
}

repositories {
    jcenter()
}

dependencies {

    implementation("com.google.guava:guava:28.0-jre")
    implementation("org.eclipse:swt:3.3.0-v3346")
    implementation("org.eclipse.swt:org.eclipse.swt.cocoa.macosx.x86_64:4.3")
    implementation("io.reactivex.rxjava3:rxjava:3.0.0-RC4")

    testImplementation("org.junit.jupiter:junit-jupiter-api:5.4.2")
    testRuntimeOnly("org.junit.jupiter:junit-jupiter-engine:5.4.2")
}

application {
    mainClassName = "io.github.andrebq.cells.microscope.swt.app.App"
    applicationDefaultJvmArgs = listOf("-XstartOnFirstThread")
}

val test by tasks.getting(Test::class) {
    useJUnitPlatform()
}
