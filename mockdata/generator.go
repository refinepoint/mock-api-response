package mockdata

import (
	"errors"
	"fmt"
	"log"
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

func generateFieldMap(fields []string, additionalParams map[string]string) map[string]string {
	// If additionalParams is nil, initialize it as an empty map
	if additionalParams == nil {
		additionalParams = make(map[string]string)
	}

	if len(fields) == 0 {
		return additionalParams
	}
	// Populate the additionalParams with the fields, assigning empty strings if not already present
	for _, field := range fields {
		// If the field does not exist, assign an empty string to it
		// Only add the field if it doesn't already exist and the field is not empty
		if field != "" {
			if _, exists := additionalParams[field]; !exists {
				additionalParams[field] = "" // Add the field with an empty value (or modify if necessary)
			}
		}
	}

	// Return the modified additionalParams
	return additionalParams
}
func GenerateCustomJSON(size int, fields []string, additionalParams map[string]string) ([]map[string]interface{}, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	// Generate the fieldMap using provided fields and params
	fieldMap := generateFieldMap(fields, additionalParams)
	// Check if fieldMap is nil or empty
	if len(fieldMap) == 0 || isMapEmpty(fieldMap) {
		log.Println("fieldMap is either nil, empty, or contains only empty values, generating dummy JSON...")
		return generateDummyJSON(size)
	}
	//	return generateDummyJSON(size)
	var result []map[string]interface{}
	for i := 0; i < size; i++ {
		// Initialize a new map for the JSON object
		obj := make(map[string]interface{})

		// Populate the map with generated values based on the fieldMap
		for field, value := range fieldMap {
			// Call generateValueBasedOnField to determine the value for each field
			// Pass the field and existing value from fieldMap
			value, err := generateValueBasedOnField(field, value)

			if err != nil {
				return nil, err
			}
			// Set the value in the map
			obj[field] = value
		}

		// Add the generated object to the result slice
		result = append(result, obj)
	}

	return result, nil
}

func isMapEmpty(fieldMap map[string]string) bool {
	// If the map has exactly one key and that key is an empty string, it's considered empty
	if len(fieldMap) == 1 {
		for key := range fieldMap {
			if key == "" {
				return true
			}
		}
	}

	// If the map has no keys, it's considered empty
	return len(fieldMap) == 0
}

// generateValueBasedOnFieldAndType generates a value based on the field and its specified type
func generateValueBasedOnFieldAndType(fieldType string) (interface{}, error) {
	// Implement your type-based value generation logic here
	fieldType = strings.ToLower(fieldType)
	switch fieldType {
	case "uuid":
		return uuid.NewString(), nil
	case "uid":
		return uuid.NewString(), nil
	case "phone":
		return generateRandomPhone(), nil
	case "number", "integer":
		return rand.Intn(100) + 100, nil
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
	names := []string{"John", "Alice", "Bob", "Sara", "Charlie", "David", "Eve", "Sophia", "James", "Lily",
		"Michael", "Anna", "Luke", "Grace", "Oliver", "Emily", "Noah", "Amelia", "Jack", "Ella", "Lucas"}

	// Randomly select one or two words from the list
	firstName := names[rand.Intn(len(names))]
	secondName := names[rand.Intn(len(names))]
	return firstName + " " + secondName
}

// Function to generate a random address
func generateRandomAddress() string {
	addresses := []string{
		"123 Elm St", "456 Oak Ave", "789 Pine Rd", "321 Maple Blvd", "654 Birch Ln",
		"1012 Cherry Ln", "2023 Walnut Dr", "4050 Pinewood Rd", "7075 Cedar Park", "1001 Sunset Blvd",
	}

	return addresses[rand.Intn(len(addresses))]
}

// Function to generate a random contact (phone number or email)
func generateRandomEmail() string {
	return generateRandomName() + "@refinepoint.com"
}
func generateRandomPhone() string {
	rand.Seed(time.Now().UnixNano())

	// Randomly select a country code (example: +1 for USA, +44 for UK)
	countryCodes := []string{"+1", "+44", "+91", "+61", "+49", "+33", "+81", "+55"}
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
