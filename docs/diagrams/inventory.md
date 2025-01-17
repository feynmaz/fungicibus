```mermaid
sequenceDiagram
    actor A as Admin
    participant F as Frontend
    participant B as Backend/DB

    A->>F: Opens admin panel
    F->>B: GET /items
    B-->>F: Return list of items
    F-->>A: Display items list

    A->>F: Enter product name & click Search
    F->>B: GET /items/{name}

    alt Product exists
        B-->>F: Return product (200 OK)
        F-->>A: Show edit page
        A->>F: Modify product details
        F->>B: PUT /items/{name}
        B-->>F: Success response
        F-->>A: Show success message
    else Product not found
        B-->>F: Return 404
        F-->>A: Show "Not found. Create new?"
        
        alt Admin realizes typo
            A->>F: Enter correct name
            F->>B: GET /items/{name}
            B-->>F: Return product (200 OK)
            F-->>A: Show edit page
        else Admin confirms creation
            A->>F: Click "Yes, create new"
            F-->>A: Show create product page
            A->>F: Fill product details
            F->>B: POST /items
            B-->>F: Success response
            F-->>A: Show success message
        end
    end

```