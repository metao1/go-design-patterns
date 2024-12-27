## **1. Decorator Pattern**

The **Decorator Pattern** dynamically adds behavior to an object without modifying its structure. This is ideal for enhancing functionality in a flexible, reusable manner.

### **Features of the Example**:
- **Base Stream Implementation:**
    - A `StringInputStream` for simple string-based input.
- **Decorators:**
    - `JsonFileStream`: Validates and parses JSON data.
    - `EncryptionDecorator`: Encrypts and decrypts data.
    - `FileDataStream`: Reads and writes file data.

### **Example Use Case**:
Enhance an input stream with encryption, logging, and validation without modifying the original implementation.
