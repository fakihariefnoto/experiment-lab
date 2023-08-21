Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying their concrete classes.

*Aplicabillity*

 Use the Abstract Factory when your code needs to work with various families of related products, but you don’t want it to depend on the concrete classes of those products—they might be unknown beforehand or you simply want to allow for future extensibility.

 The Abstract Factory provides you with an interface for creating objects from each class of the product family. As long as your code creates objects via this interface, you don’t have to worry about creating the wrong variant of a product which doesn’t match the products already created by your app.

 Consider implementing the Abstract Factory when you have a class with a set of Factory Methods that blur its primary responsibility.

 In a well-designed program each class is responsible only for one thing. When a class deals with multiple product types, it may be worth extracting its factory methods into a stand-alone factory class or a full-blown Abstract Factory implementation.

  
*How to Implement*

Map out a matrix of distinct product types versus variants of these products.

Declare abstract product interfaces for all product types. Then make all concrete product classes implement these interfaces.

Declare the abstract factory interface with a set of creation methods for all abstract products.

Implement a set of concrete factory classes, one for each product variant.

Create factory initialization code somewhere in the app. It should instantiate one of the concrete factory classes, depending on the application configuration or the current environment. Pass this factory object to all classes that construct products.

Scan through the code and find all direct calls to product constructors. Replace them with calls to the appropriate creation method on the factory object.

*Pros and Cons*
(+)
You can be sure that the products you’re getting from a factory are compatible with each other.
You avoid tight coupling between concrete products and client code.
Single Responsibility Principle. You can extract the product creation code into one place, making the code easier to support.
Open/Closed Principle. You can introduce new variants of products without breaking existing client code.

(-)
The code may become more complicated than it should be, since a lot of new interfaces and classes are introduced along with the pattern.