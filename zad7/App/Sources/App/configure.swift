import NIOSSL
import Fluent
import FluentSQLiteDriver
import Leaf
import Vapor

/// configures your application
func configure(_ app: Application) async throws {
    // uncomment to serve files from /Public folder
    // app.middleware.use(FileMiddleware(publicDirectory: app.directory.publicDirectory))

    app.databases.use(DatabaseConfigurationFactory.sqlite(.file("db.sqlite")), as: .sqlite)

    app.migrations.add(CreateTodo())
    app.migrations.add(CreateProduct())

    app.views.use(.leaf)

    try await app.autoMigrate()

    if try await Product.query(on: app.db).count() == 0 {
        let sampleProducts = [
            Product(name: "Apple", price: 1.49),
            Product(name: "Banana", price: 0.89),
            Product(name: "Coffee", price: 4.99)
        ]

        for product in sampleProducts {
            try await product.save(on: app.db)
        }

        app.logger.info("Seeded sample products")
    }

    // register routes
    try routes(app)
}
