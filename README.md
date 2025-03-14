## API Documentation  

###  GET `/test`  
#### Description  
Returns the current time in IST (Asia/Kolkata).  

#### Response Example  
```json
{
  "message": "Success 2025-03-15 18:45:30 IST"
}
```

---

###  POST `/parex_test`  
#### Description  
Uploads a file and stores it in `./tmp/`.  

#### Request (Multipart Form-Data)  
| Field | Type | Description |
|--------|------|-------------|
| `file` | File | Any file to upload |

#### Response Example  
```json
{
  "message": "File uploaded successfully",
  "file_path": "./tmp/sample.pdf"
}
```

---

###  POST `/parex_process`  
#### Description  
Processes the uploaded file using `lib.Explore`.  

#### Request (Multipart Form-Data)  
| Field   | Type   | Description                    |
|---------|--------|--------------------------------|
| `file`  | File   | File to process (stored in `./tmp/`) |
| `offset` | Text  | Offset value (e.g., `1024`)   |
| `level` | Text   | Processing level (`0-3`)      |

#### Response Example (Success)  
```json
{
  "message": "Processing completed successfully"
}
```

#### Response Example (Error)  
```json
{
  "error": "Error processing file",
  "details": "specific error message"
}
```

---

## Testing the API with Postman  

###  Test `/test` Route  
- Method: GET  
- URL: `http://localhost:8080/test`  
- Expected Response: `{ "message": "Success {IST Time}" }`  

---

###  Test `/parex_test` (File Upload Route)  
1. Open Postman  
2. Method: `POST`  
3. URL: `http://localhost:8080/parex_test`  
4. Go to Body → Select "form-data"  
5. Key: `file` → Type: `File` → Upload a file  
6. Click "Send"  
7. Expected Response:  
```json
{
  "message": "File uploaded successfully",
  "file_path": "./tmp/sample.pdf"
}
```

---

###  Test `/parex_process` (File Processing Route)  
1. Open Postman  
2. Method: `POST`  
3. URL: `http://localhost:8080/parex_process`  
4. Go to Body → Select "form-data"  
5. Add the following fields:  
   - Key: `file` → Type: `File` → Upload a file  
   - Key: `offset` → Type: `Text` → Example: `1024`  
   - Key: `level` → Type: `Text` → Example: `2`  
6. Click "Send"  
7. Expected Response:  
```json
{
  "message": "Processing completed successfully"
}
```

---
