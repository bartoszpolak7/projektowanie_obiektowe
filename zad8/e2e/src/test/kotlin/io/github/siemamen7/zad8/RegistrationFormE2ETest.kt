package io.github.siemamen7.zad8

import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.Assertions.assertTrue
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import org.openqa.selenium.By
import org.openqa.selenium.JavascriptExecutor
import org.openqa.selenium.WebDriver
import org.openqa.selenium.chrome.ChromeDriver
import org.openqa.selenium.chrome.ChromeOptions
import org.openqa.selenium.support.ui.ExpectedConditions
import org.openqa.selenium.support.ui.WebDriverWait
import java.time.Duration

class RegistrationFormE2ETest {
    private lateinit var driver: WebDriver
    private lateinit var wait: WebDriverWait

    @BeforeEach
    fun setUp() {
        val options = ChromeOptions().apply {
            setBinary("C:\\Program Files\\BraveSoftware\\Brave-Browser\\Application\\brave.exe")
            addArguments("--headless=new", "--no-sandbox", "--disable-dev-shm-usage")
        }
        driver = ChromeDriver(options)
        wait = WebDriverWait(driver, Duration.ofSeconds(10))
    }

    @AfterEach
    fun tearDown() {
        if (this::driver.isInitialized) {
            driver.quit()
        }
    }

    @Test
    fun shouldRequireEmailAndPasswordBeforeSubmitting() {
        openRegistrationPage()

        driver.findElement(By.cssSelector("button")).click()

        assertTrue(fieldValidity("email", "valueMissing"), "Email should be marked as required")
        assertTrue(fieldValidity("password", "valueMissing"), "Password should be marked as required")
        assertTrue(driver.currentUrl?.endsWith("/register") == true, "The form should not submit when fields are empty")
    }

    @Test
    fun shouldPreventSubmissionForInvalidEmailFormat() {
        openRegistrationPage()

        driver.findElement(By.id("email")).sendKeys("invalid-email")
        driver.findElement(By.id("password")).sendKeys("StrongPassword123")
        driver.findElement(By.cssSelector("button")).click()

        assertTrue(fieldValidity("email", "typeMismatch"), "Invalid email format should be rejected")
        assertTrue(driver.currentUrl?.endsWith("/register") == true, "The form should not submit for invalid email")
    }

    private fun openRegistrationPage() {
        val baseUrl = System.getenv("ZAD8_FRONTEND_URL") ?: "http://localhost:5173/register"
        driver.get(baseUrl)
        wait.until(ExpectedConditions.visibilityOfElementLocated(By.tagName("h1")))
        wait.until(ExpectedConditions.visibilityOfElementLocated(By.id("email")))
        wait.until(ExpectedConditions.visibilityOfElementLocated(By.id("password")))
    }

    private fun fieldValidity(fieldId: String, validityProperty: String): Boolean {
        val js = driver as JavascriptExecutor
        return js.executeScript("return document.getElementById('$fieldId').validity.$validityProperty;") as Boolean
    }
}
