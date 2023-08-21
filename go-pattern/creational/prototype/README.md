Prototype is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.

*Applicability* 

*- Use the Prototype pattern when your code shouldn’t depend on the concrete classes of objects that you need to copy. -*

 This happens a lot when your code works with objects passed to you from 3rd-party code via some interface. The concrete classes of these objects are unknown, and you couldn’t depend on them even if you wanted to.

The Prototype pattern provides the client code with a general interface for working with all objects that support cloning. This interface makes the client code independent from the concrete classes of objects that it clones.

*- Use the pattern when you want to reduce the number of subclasses that only differ in the way they initialize their respective objects. -*

 Suppose you have a complex class that requires a laborious configuration before it can be used. There are several common ways to configure this class, and this code is scattered through your app. To reduce the duplication, you create several subclasses and put every common configuration code into their constructors. You solved the duplication problem, but now you have lots of dummy subclasses.

The Prototype pattern lets you use a set of pre-built objects configured in various ways as prototypes. Instead of instantiating a subclass that matches some configuration, the client can simply look for an appropriate prototype and clone it.