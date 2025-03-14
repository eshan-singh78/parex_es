## API Documentation   

### 1. GET `/test`  
#### Description  
Returns the current time in IST (Asia/Kolkata).  

#### Response Example  
```json
{
  "message": "Success 2025-03-15 18:45:30 IST"
}
```

---

### 2. POST `/parex_test`  
#### Description  
Uploads a file and stores it in `./tmp/`.  

#### Request (Multipart Form-Data)  
| Field  | Type | Description |
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

### 3. POST `/parex_process_v2`  
#### Description  
Processes the uploaded file using `lib.Explore` and returns a list of extracted filenames.  

#### Request (Multipart Form-Data)  
| Field   | Type   | Description                      |
|---------|--------|----------------------------------|
| `file`  | File   | File to process (stored in `./tmp/`) |
| `offset` | Text  | Offset value (e.g., `1024`)     |
| `level` | Text   | Processing level (`0-3`)        |

#### Response Example (Success)  
```json
{
  "message": "Processing completed successfully",
  "filenames": [
    "WOQVU6AIMB.png",
    "UUU7Y77U4O.docx",
    "2YQSXPBI5F.txt"
  ]
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

### 1. Test `/test` Route  
- Method: GET  
- URL: `http://localhost:8080/test`  
- Expected Response: `{ "message": "Success {IST Time}" }`  

---

### 2. Test `/parex_test` (File Upload Route)  
1. Open Postman  
2. Method: `POST`  
3. URL: `http://localhost:8080/parex_test`  
4. Go to Body → Select "form-data"  
5. Add the following field:  
   - Key: `file` → Type: `File` → Upload any file  
6. Click "Send"  
7. Expected Response:  
```json
{
  "message": "File uploaded successfully",
  "file_path": "./tmp/sample.pdf"
}
```

---

### 3. Test `/parex_process_v2` (File Processing with Filenames Extraction)  
1. Open Postman  
2. Method: `POST`  
3. URL: `http://localhost:8080/parex_process_v2`  
4. Go to Body → Select "form-data"  
5. Add the following fields:  
   - Key: `file` → Type: `File` → Upload a file  
   - Key: `offset` → Type: `Text` → Example: `1024`  
   - Key: `level` → Type: `Text` → Example: `2`  
6. Click "Send"  
7. Expected Response (if successful):  
```json
{
  "message": "Processing completed successfully",
  "filenames": [
    "WOQVU6AIMB.png",
    "UUU7Y77U4O.docx",
    "2YQSXPBI5F.txt"
  ]
}
```
8. Expected Response (if error occurs):  
```json
{
  "error": "Error processing file",
  "details": "specific error message"
}
```

---
