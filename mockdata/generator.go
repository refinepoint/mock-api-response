package mockdata

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var words = []string{
	"apple", "banana", "orange", "grape", "mango", "strawberry", "peach", "kiwi", "blueberry", "pineapple",
	"mountain", "ocean", "river", "forest", "desert", "lake", "cloud", "sky", "sun", "moon",
	"happy", "sad", "angry", "excited", "calm", "nervous", "joyful", "peaceful", "relaxed", "bored",
	"dog", "cat", "rabbit", "lion", "tiger", "elephant", "zebra", "giraffe", "horse", "kangaroo",
}

// Predefined sets of words for sentence generation
var subjects = []string{"The cat", "A dog", "The man", "She", "John", "The teacher", "The woman", "The child", "Our team", "They"}
var verbs = []string{"eats", "plays with", "writes", "runs", "sings", "loves", "is", "creates", "discovers", "believes"}
var adjectives = []string{"happy", "fast", "lazy", "beautiful", "angry", "exciting", "interesting", "smart", "quiet", "wonderful"}
var objects = []string{"the ball", "a book", "the piano", "the story", "a new car", "the computer", "an idea", "the project", "the song", "a dream"}
var adverbs = []string{"quickly", "carefully", "happily", "loudly", "gracefully", "excitedly", "eagerly", "clumsily", "calmly", "seriously"}
var conjunctions = []string{"and", "but", "so", "because", "although"}

// UUID Fields
var uuidFields = []string{
	"uuid", "uid", "id", "transactionId", "userId", "sessionId", "orderId", "productId", "customerId", "invoiceId",
	"paymentId", "bookingId", "messageId", "requestId", "documentId", "recordId", "subscriptionId", "assetId",
	"fileId", "eventId", "deviceId", "jobId", "processId", "customerUuid", "orderUuid", "transactionUuid", "paymentUuid",
	"sessionUuid", "invoiceUuid", "fileUuid", "messageUuid", "deviceUuid", "bookingUuid", "requestUuid", "documentUuid",
	"eventUuid", "userUuid", "productUuid", "transactionKey", "userKey", "orderKey", "uniqueId", "entityId",
}

// DateTime Fields
var datetimeFields = []string{
	"date", "datetime", "timestamp", "createdAt", "updatedAt", "lastLogin", "expiryDate", "startDate", "endDate", "birthDate",
	"orderDate", "completionTime", "scheduledTime", "dueDate", "publishedAt", "deletedAt", "checkInTime", "checkOutTime",
	"eventDate", "approvalDate", "processingTime", "modificationDate", "lastUpdated", "lastAccessed", "createdOn", "expiredAt",
	"deliveryDate", "dispatchTime", "arrivalTime", "joinDate", "lastModified", "confirmedAt", "reviewDate", "cancellationDate",
	"invoiceDate", "paymentDate", "subscriptionDate", "renewalDate", "confirmationDate", "activationDate", "lastSeen",
	"responseTime", "entryTime", "exitTime", "serviceDate", "contractDate", "appointmentDate",
}

// Numeric Fields
var numericFields = []string{
	"integer", "number", "quantity", "price", "discount", "totalAmount", "score", "rank", "level", "age",
	"orderNumber", "invoiceNumber", "balance", "rating", "latitude", "longitude", "transactionAmount", "phoneNumber",
	"duration", "limit", "attempts", "progressPercentage", "totalQuantity", "totalPrice", "unitPrice", "balanceAmount",
	"weight", "height", "width", "depth", "distance", "timeSpent", "amountPaid", "totalSpent", "totalIncome",
	"dailyLimit", "totalValue", "debtAmount", "fineAmount", "taxAmount", "profit", "lossAmount", "surplusAmount",
	"orderCount", "transactionCount", "remainingAmount", "totalDiscount", "deliveryFee",
}

// Name Fields
var nameFields = []string{
	"firstName", "lastName", "middleName", "fullName", "nickName", "contactName", "userName", "displayName", "name",
	"companyName", "brandName", "shopName", "ownerName", "managerName", "employeeName", "productName", "serviceName",
	"customerName", "partnerName", "supplierName", "vendorName", "serviceProviderName", "doctorName", "patientName",
	"teacherName", "studentName", "staffName", "authorName", "editorName", "publisherName", "directorName", "designerName",
	"coAuthorName", "coEditorName", "teamMemberName", "projectManagerName", "supportAgentName", "serviceRepName", "assistantName",
	"contactPerson", "legalName", "socialName", "brandOwnerName", "founderName", "ownerFullName", "recipientName",
}

// Address Fields
var addressFields = []string{
	"address", "street", "city", "state", "zipCode", "country", "postalCode", "province", "buildingNumber", "streetName",
	"apartment", "suite", "neighborhood", "district", "region", "location", "addressLine1", "addressLine2", "fullAddress",
	"billingAddress", "shippingAddress", "deliveryAddress", "homeAddress", "officeAddress", "permanentAddress", "temporaryAddress",
	"addressDetails", "cityName", "stateName", "countryName", "addressCode",
}

// Contact Fields
var emailFields = []string{
	"email", "contactEmail", "primaryEmail",
	"secondaryEmail", "workEmail", "personalEmail", "emergencyContact", "supportEmail",
	"contactPerson", "linkedinProfile", "twitterHandle",
	"facebookProfile", "instagramHandle",
}
var phoneFields = []string{
	"phone", "phoneNumber", "contactNumber", "mobile", "homePhone", "officePhone", "fax", "emergencyContact", "socialSecurityNumber", "taxNumber", "customerPhone",
}

// Status Fields
var statusFields = []string{
	"status", "state", "currentStatus", "orderStatus", "paymentStatus", "accountStatus", "userStatus", "subscriptionStatus",
	"processStatus", "jobStatus", "taskStatus", "projectStatus", "ticketStatus", "documentStatus", "invoiceStatus", "shipmentStatus",
	"serviceStatus", "orderState", "accountState", "approvalStatus", "confirmationStatus", "activationStatus", "deactivationStatus",
	"cancellationStatus", "completionStatus", "deliveryStatus", "validationStatus", "approvalState", "reviewStatus", "feedbackStatus",
	"finalStatus", "userState", "notificationStatus", "emailStatus", "phoneStatus", "paymentState", "requestStatus",
}
var longText = []string{
	"description", "desc", "des", "longtext", "paragraph", "para",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Function to generate dummy data of a specified size
// Updated generateDummyJSON to return []map[string]interface{}
func generateDummyJSON(size int) ([]map[string]interface{}, error) {
	dummyData := make([]map[string]interface{}, size)

	for i := 0; i < size; i++ {
		dummyData[i] = map[string]interface{}{
			"id":    strconv.Itoa(i + 1),
			"name":  "Item " + strconv.Itoa(i+1),
			"price": rand.Float64() * 100, // Generate a random price between 0 and 100
			"desc":  "This is a description for item " + strconv.Itoa(i+1),
		}
	}

	return dummyData, nil
}

//

func generateFieldMap(fields []string, additionalParams map[string]string) map[string]interface{} {
	// If additionalParams is nil, initialize it as an empty map
	if additionalParams == nil {
		additionalParams = make(map[string]string)
	}

	fieldMap := make(map[string]interface{})

	if len(fields) == 0 {
		// Copy all values from additionalParams as is
		for key, value := range additionalParams {
			fieldMap[key] = value
		}
		return fieldMap
	}

	// Populate the fieldMap with fields and additionalParams
	for _, field := range fields {
		if field != "" {
			addNestedField(fieldMap, field, "")
		}
	}
	for key, value := range additionalParams {
		addNestedField(fieldMap, key, value)

	}
	return fieldMap
}

// Function to add nested fields into a map
func addNestedField(fieldMap map[string]interface{}, field string, value interface{}) {
	parts := strings.Split(field, ".")
	current := fieldMap

	for i, part := range parts {
		// If it's the last part, set the value
		if i == len(parts)-1 {
			current[part] = value
		} else {
			// If the part doesn't exist or isn't a map, create a new map
			if _, exists := current[part]; !exists {
				current[part] = make(map[string]interface{})
			}

			// Move deeper into the nested map
			if next, ok := current[part].(map[string]interface{}); ok {
				current = next
			} else {
				// Handle case where the key already exists but isn't a map
				current[part] = make(map[string]interface{})
				current = current[part].(map[string]interface{})
			}
		}
	}
}

func GenerateCustomJSON(size int, fields []string, additionalParams map[string]string, encloseInArray bool) (interface{}, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	// Generate the fieldMap using provided fields and params
	fieldMap := generateFieldMap(fields, additionalParams)
	// Check if fieldMap is nil or empty
	if len(fieldMap) == 0 || isMapEmpty(fieldMap) {
		return generateDummyJSON(size)
	}

	if size == 1 && !encloseInArray {
		// Generate a single JSON object when size is 1
		obj := make(map[string]interface{})
		populateNestedFields(obj, fieldMap)
		return obj, nil
	}

	var result []map[string]interface{}
	for i := 0; i < size; i++ {
		// Generate JSON object with nested fields
		obj := make(map[string]interface{})
		populateNestedFields(obj, fieldMap)

		// Add the generated object to the result slice
		result = append(result, obj)
	}

	return result, nil
}

// Function to populate nested fields with values
func populateNestedFields(obj map[string]interface{}, fieldMap map[string]interface{}) {
	for key, value := range fieldMap {
		switch v := value.(type) {
		case map[string]interface{}:
			// Create a nested map if necessary
			if _, exists := obj[key]; !exists {
				obj[key] = make(map[string]interface{})
			}
			// Recursively populate nested fields
			if nestedObj, ok := obj[key].(map[string]interface{}); ok {
				populateNestedFields(nestedObj, v)
			}
		default:
			// Generate value based on the field
			strValue, _ := v.(string)

			generatedValue, err := generateValueBasedOnField(key, strValue)
			if err == nil {
				obj[key] = generatedValue
			}
		}
	}
}

func isMapEmpty(fieldMap map[string]interface{}) bool {
	for key, value := range fieldMap {
		// Check if key or value is empty
		if key == "" {
			return true
		}

		// Check for nested maps recursively
		if nestedMap, ok := value.(map[string]interface{}); ok {
			if isMapEmpty(nestedMap) {
				return true
			}
		}
	}

	return len(fieldMap) == 0
}

// generateValueBasedOnFieldAndType generates a value based on the field and its specified type
func generateValueBasedOnFieldAndType(fieldType string) (interface{}, error) {
	// Implement your type-based value generation logic here
	fieldType = strings.ToLower(fieldType)
	switch fieldType {
	case "uuid", "uid", "id":
		return uuid.NewString(), nil
	case "phone":
		return generateRandomPhone(), nil
	case "number", "integer":
		return rand.Intn(100) + 100, nil
	case "date", "datetime", "time":
		return generateRandomDatetime(), nil
	case "email", "mail":
		return generateRandomEmail(), nil
	case "long":
		return rand.Int63(), nil
	case "float", "decimal":

		roundedDecimal := float64(int(rand.Float64()*10000)) / 10000
		// Return the final value with an offset of +10
		return roundedDecimal + 10, nil
	case "text":
		return generateRandomWordOrTwo(), nil
	case "boolean":
		// Return "true" for 0, "false" for 1
		if rand.Intn(2) == 0 {
			return "true", nil
		}
		return "false", nil
	default:
		return generateRandomWordOrTwo(), nil
	}
}

func generateValueBasedOnField(field string, fieldType string) (interface{}, error) {
	if fieldType != "" {
		return generateValueBasedOnFieldAndType(fieldType)
	}

	field = normalizeFieldName(field)

	if strings.Contains(field, "id") {
		return uuid.NewString(), nil
	}
	switch {

	case contains(datetimeFields, field):
		return generateRandomDatetime(), nil
	case contains(numericFields, field):
		return generateRandomNumber(), nil
	case contains(nameFields, field):
		return generateRandomName(), nil
	case contains(addressFields, field):
		return generateRandomAddress(), nil
	case contains(phoneFields, field):
		return generateRandomPhone(), nil
	case contains(emailFields, field):
		return generateRandomEmail(), nil
	case contains(statusFields, field):
		return generateRandomStatus(), nil
	case contains(uuidFields, field):
		return uuid.NewString(), nil
	case contains(longText, field):
		return generateSentence(), nil
	default:
		return generateRandomWordOrTwo(), nil

	}
}

func contains(list []string, key string) bool {
	for _, item := range list {
		if normalizeFieldName(item) == key {
			return true
		}
	}
	return false
}

func normalizeFieldName(field string) string {
	return strings.ToLower(field)
}

// Function to generate a random name (single word or two words)
func generateRandomName() string {
	names := []string{
		"John", "Alice", "Bob", "Sara", "Charlie", "David", "Eve", "Sophia", "James", "Lily",
		"Michael", "Anna", "Luke", "Grace", "Oliver", "Emily", "Noah", "Amelia", "Jack", "Ella", "Lucas",

		// Indian Names
		"Aarav", "Ishaan", "Vivaan", "Aditya", "Krishna", "Saanvi", "Ananya", "Diya", "Rohan", "Aadhya",
		"Priya", "Kavya", "Raj", "Anil", "Simran", "Neha", "Amit", "Pooja", "Rani", "Shivani", "Mohit",
		"Varun", "Shreya", "Manish", "Geeta", "Madhavi", "Ravi", "Meera", "Ritika", "Sunil", "Pankaj",

		// Japanese Names
		"Hiroshi", "Yuki", "Haruto", "Sakura", "Takumi", "Yui", "Riku", "Miyu", "Ren", "Aiko",
		"Kaito", "Emi", "Taro", "Hina", "Yuuto", "Rika", "Satoshi", "Aya", "Shun", "Yoshiko", "Kenta",
		"Sayaka", "Sho", "Rena", "Kumi", "Koji", "Mai", "Daiki", "Nanami", "Shiori", "Eiko",

		// Tamil Names
		"Arjun", "Karthik", "Vijay", "Rajesh", "Sivakumar", "Anjali", "Kavitha", "Sneha", "Ravi", "Nisha",
		"Ganesh", "Priya", "Saravanan", "Lavanya", "Pavithra", "Hari", "Lakshmi", "Vignesh", "Selvi", "Sundar",
		"Mani", "Sindhu", "Krithika", "Suresh", "Madhavi", "Meenakshi", "Gokul", "Pooja", "Karthika", "Bhavani",

		// Chinese Names
		"Li Wei", "Wang Wei", "Li Na", "Zhang Wei", "Liu Yang", "Chen Xi", "Wang Fang", "Zhao Li", "Li Jing", "Wang Jun",
		"Li Mei", "Yang Li", "Zhang Jing", "Liu Feng", "Chen Li", "Li Hong", "Zhao Yan", "Wang Lei", "Chen Wei", "Li Yun",
		"Yang Mei", "Liu Jin", "Zhao Hong", "Wang Hui", "Zhang Fang", "Li Ping", "Chen Jing", "Yang Wei", "Wang Xue", "Li Fang",
	}

	// Randomly select one or two words from the list
	firstName := names[rand.Intn(140)]
	secondName := names[rand.Intn(140)]
	return firstName + " " + secondName
}

// Function to generate a random address
func generateRandomAddress() string {

	addresses := []string{
		"123 Elm St", "456 Oak Ave", "789 Pine Rd", "321 Maple Blvd", "654 Birch Ln",
		"1012 Cherry Ln", "2023 Walnut Dr", "4050 Pinewood Rd", "7075 Cedar Park", "1001 Sunset Blvd",

		// India Addresses
		"1 MG Road, Bangalore, Karnataka, India", "10 Connaught Place, New Delhi, India", "33 Nehru Place, Delhi, India",
		"58 Bandra West, Mumbai, Maharashtra, India", "11 Park Street, Kolkata, West Bengal, India",
		"45 Marine Drive, Mumbai, Maharashtra, India", "12 Brigade Road, Bangalore, Karnataka, India", "7 Juhu Beach, Mumbai, India",
		"21 Rajaji Salai, Chennai, Tamil Nadu, India", "101 Lower Tank Bund, Hyderabad, Telangana, India",

		// UK Addresses
		"221B Baker Street, London, UK", "10 Downing Street, London, UK", "5 The Crescent, Birmingham, UK",
		"4 St Andrew's Square, Edinburgh, Scotland", "12 Oxford Street, Manchester, UK", "56 King Street, London, UK",
		"20 Abbey Road, London, UK", "8 Victoria Street, London, UK", "34 High Street, Glasgow, Scotland", "7 Princes Street, Edinburgh, UK",

		// USA Addresses
		"1600 Pennsylvania Ave NW, Washington, DC, USA", "742 Evergreen Terrace, Springfield, USA", "101 California St, San Francisco, CA, USA",
		"350 5th Ave, New York, NY, USA", "1000 16th St NW, Washington, DC, USA", "4750 Drexel Dr, Los Angeles, CA, USA",
		"200 E 74th St, New York, NY, USA", "12345 Sunset Blvd, Los Angeles, CA, USA", "789 Beach Ave, Miami, FL, USA",
		"2020 W Broadway, Vancouver, BC, Canada",

		// Australia Addresses
		"100 Collins St, Melbourne, VIC, Australia", "123 Pitt St, Sydney, NSW, Australia", "22 George St, Brisbane, QLD, Australia",
		"88 Queen St, Melbourne, VIC, Australia", "43 Bondi Beach, Sydney, NSW, Australia", "150 Swanston St, Melbourne, VIC, Australia",
		"11 Southbank, Sydney, NSW, Australia", "28 King St, Melbourne, VIC, Australia", "9 Circular Quay, Sydney, NSW, Australia",
		"5 Phillip St, Sydney, NSW, Australia",

		// Canada Addresses
		"1 Yonge St, Toronto, ON, Canada", "123 Main St, Vancouver, BC, Canada", "56 Bay St, Toronto, ON, Canada",
		"44 Spadina Ave, Toronto, ON, Canada", "12 Richmond St, Ottawa, ON, Canada", "888 St. Denis St, Montreal, QC, Canada",
		"555 King St W, Toronto, ON, Canada", "87 Wellington St W, Toronto, ON, Canada", "1200 Steeles Ave, Toronto, ON, Canada",
		"1050 Rue St. Paul E, Montreal, QC, Canada",

		// European Addresses
		"5 Rue de la Paix, Paris, France", "30 Via Roma, Rome, Italy", "7 Bahnhofstrasse, Zurich, Switzerland",
		"25 Oxford St, London, UK", "21 Elizabetta St, Milan, Italy", "10 Av. des Champs-Élysées, Paris, France",
		"15 Puerta del Sol, Madrid, Spain", "42 Königsstraße, Berlin, Germany", "9 Les Champs-Élysées, Paris, France", "23 Rue des Archers, Brussels, Belgium",

		// Middle East Addresses
		"6 Sheikh Zayed Road, Dubai, UAE", "25 Corniche Rd, Abu Dhabi, UAE", "120 Jeddah St, Riyadh, Saudi Arabia",
		"8 Al Haram St, Cairo, Egypt", "19 Al Faisaliyah, Riyadh, Saudi Arabia", "3 Khaled Bin Waleed St, Sharjah, UAE",
		"77 Beirut Street, Lebanon", "45 Al Ahram St, Cairo, Egypt", "22 Mohammed Bin Rashid Blvd, Dubai, UAE", "112 Dubai Marina, Dubai, UAE",

		// African Addresses
		"12 Nelson Mandela Dr, Pretoria, South Africa", "34 Victoria Island, Lagos, Nigeria", "5 Harare Rd, Harare, Zimbabwe",
		"77 Anwar St, Nairobi, Kenya", "150 Nkrumah Ave, Accra, Ghana", "8 Tanganyika Rd, Dar Es Salaam, Tanzania", "10 Pali Hill, Mumbai, India",
		"6 Independence Ave, Windhoek, Namibia", "21 Kombo Beach, Banjul, Gambia", "42 Cairo St, Cairo, Egypt",
	}

	return addresses[rand.Intn(len(addresses))]
}

// Function to generate a random contact (phone number or email)
func generateRandomEmail() string {
	return strings.ReplaceAll(generateRandomName(), " ", ".") + "@rpoint.com"
}
func generateRandomPhone() string {
	rand.Seed(time.Now().UnixNano())

	// Randomly select a country code (example: +1 for USA, +44 for UK)
	countryCodes := []string{"+1", "+44", "+91", "+61", "+49", "+33", "+81", "+55", "+0 "}
	countryCode := countryCodes[rand.Intn(len(countryCodes))]

	// Generate a 10-digit phone number
	areaCode := rand.Intn(900) + 100          // Random area code between 100 and 999
	centralOfficeCode := rand.Intn(900) + 100 // Random central office code between 100 and 999
	lineNumber := rand.Intn(9000) + 1000      // Random line number between 1000 and 9999

	// Return the phone number in the format "+<country code> <area code>-<central office code>-<line number>"
	return fmt.Sprintf("%s %d-%d-%d", countryCode, areaCode, centralOfficeCode, lineNumber)
}

// Function to generate a random status
func generateRandomStatus() string {
	statuses := []string{
		"active", "inactive", "pending", "approved", "rejected", "completed", "processing", "shipped", "delivered", "failed",
		"cancelled", "confirmed", "successful", "unsuccessful", "onHold", "inProgress", "completed", "finalized",
	}

	return statuses[rand.Intn(len(statuses))]
}

func generateRandomDatetime() string {
	return time.Now().Add(time.Duration(rand.Intn(1000)-500) * time.Hour).Format(time.RFC3339)
}

func generateRandomNumber() int {
	return rand.Intn(10000) // Random number up to 10,000
}

func generateRandomWordOrTwo() string {
	count := rand.Intn(5) + 1 // Generates a random number between 1 and 5

	// Create a slice to store the words
	var result []string

	// Generate the random number of words
	for i := 0; i < count; i++ {
		result = append(result, words[rand.Intn(len(words))])
	}

	// Join the words with a space separator and return the result
	return strings.Join(result, " ")
}

func generateSentence() string {
	// Randomly select a sentence structure
	structure := rand.Intn(6)

	// Randomly pick sentence components
	subject := subjects[rand.Intn(len(subjects))]
	verb := verbs[rand.Intn(len(verbs))]
	object := objects[rand.Intn(len(objects))]
	adjective := adjectives[rand.Intn(len(adjectives))]
	adverb := adverbs[rand.Intn(len(adverbs))]
	conjunction := conjunctions[rand.Intn(len(conjunctions))]

	// Build sentences based on structure
	var sentence string
	switch structure {
	case 0: // Subject + Verb + Object
		sentence = fmt.Sprintf("%s %s %s.", subject, verb, object)
	case 1: // Subject + Verb + Adverb + Object
		sentence = fmt.Sprintf("%s %s %s %s.", subject, verb, adverb, object)
	case 2: // Subject + Verb + Adjective + Object
		sentence = fmt.Sprintf("%s %s %s %s.", subject, verb, adjective, object)
	case 3: // Subject + Verb + Object + Adverb
		sentence = fmt.Sprintf("%s %s %s %s.", subject, verb, object, adverb)
	case 4: // Subject + Verb + Adjective + Object + Adverb
		sentence = fmt.Sprintf("%s %s %s %s %s.", subject, verb, adjective, object, adverb)
	case 5: // Compound sentence (two parts)
		// Randomly create two parts and join them with a conjunction
		subject2 := subjects[rand.Intn(len(subjects))]
		verb2 := verbs[rand.Intn(len(verbs))]
		object2 := objects[rand.Intn(len(objects))]
		part1 := fmt.Sprintf("%s %s %s", subject, verb, object)
		part2 := fmt.Sprintf("%s %s %s", subject2, verb2, object2)
		sentence = fmt.Sprintf("%s %s %s.", part1, conjunction, part2)
	}

	return sentence
}
