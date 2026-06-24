import Fluent
import Vapor

struct ProductController: RouteCollection {
    func boot(routes: any RoutesBuilder) throws {
        let products = routes.grouped("products")

        products.get(use: index)
        products.post("create", use: create)
        products.get("new", use: newForm)
        products.get(":productID", use: show)
        products.get(":productID", "edit", use: editForm)
        products.post(":productID", "update", use: update)
        products.post(":productID", "delete", use: delete)
    }

    @Sendable
    func index(req: Request) async throws -> View {
        let products = try await Product.query(on: req.db).all()
        return try await req.view.render("products/index", ["products": products])
    }

    @Sendable
    func newForm(req: Request) async throws -> View {
        try await req.view.render("products/new")
    }

    @Sendable
    func show(req: Request) async throws -> View {
        guard let product = try await Product.find(req.parameters.get("productID"), on: req.db) else {
            throw Abort(.notFound)
        }
        return try await req.view.render("products/show", ["product": product])
    }

    @Sendable
    func editForm(req: Request) async throws -> View {
        guard let product = try await Product.find(req.parameters.get("productID"), on: req.db) else {
            throw Abort(.notFound)
        }
        return try await req.view.render("products/edit", ["product": product])
    }

    @Sendable
    func create(req: Request) async throws -> Response {
        let data = try req.content.decode(ProductForm.self)
        let product = Product(name: data.name, price: data.price)
        try await product.save(on: req.db)
        return req.redirect(to: "/products")
    }

    @Sendable
    func update(req: Request) async throws -> Response {
        guard let product = try await Product.find(req.parameters.get("productID"), on: req.db) else {
            throw Abort(.notFound)
        }
        let data = try req.content.decode(ProductForm.self)
        product.name = data.name
        product.price = data.price
        try await product.update(on: req.db)
        return req.redirect(to: "/products")
    }

    @Sendable
    func delete(req: Request) async throws -> Response {
        guard let product = try await Product.find(req.parameters.get("productID"), on: req.db) else {
            throw Abort(.notFound)
        }
        try await product.delete(on: req.db)
        return req.redirect(to: "/products")
    }
}

struct ProductForm: Content {
    let name: String
    let price: Double
}
