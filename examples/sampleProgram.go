/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"sync"

	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

var appConfigurationServiceInstance *appconfigurationv1.AppConfigurationV1

func initAndReturnSingletonInstanceWithAPIKey(authToken string, guid string, region string) *appconfigurationv1.AppConfigurationV1 {

	instanceURL := "https://" + region + ".apprapp.cloud.ibm.com/apprapp/feature/v1/instances/" + guid
	var once sync.Once
	if appConfigurationServiceInstance == nil {
		once.Do(func() {
			if appConfigurationServiceInstance == nil {
				authenticator := &core.IamAuthenticator{
					ApiKey: authToken,
				}
				options := &appconfigurationv1.AppConfigurationV1Options{
					Authenticator: authenticator,
					URL:           instanceURL,
				}
				var error error
				appConfigurationServiceInstance, error = appconfigurationv1.NewAppConfigurationV1(options)
				if error != nil {
					fmt.Println("Error: " + error.Error())
					return
				}
			}
		})
	}
	return appConfigurationServiceInstance
}

func initAndReturnSingletonInstanceWithBearerToken(authToken string, guid string, region string) *appconfigurationv1.AppConfigurationV1 {

	instanceURL := "https://" + region + ".apprapp.cloud.ibm.com/apprapp/feature/v1/instances/" + guid
	var once sync.Once
	if appConfigurationServiceInstance == nil {
		once.Do(func() {
			if appConfigurationServiceInstance == nil {
				authenticator := &core.BearerTokenAuthenticator{
					BearerToken: authToken,
				}
				options := &appconfigurationv1.AppConfigurationV1Options{
					Authenticator: authenticator,
					URL:           instanceURL,
				}
				var error error
				appConfigurationServiceInstance, error = appconfigurationv1.NewAppConfigurationV1(options)
				if error != nil {
					fmt.Println("Error: " + error.Error())
					return
				}
			}
		})
	}
	return appConfigurationServiceInstance
}

func main() {

	authToken := "<authToken>"
	guid := "<guid>"
	region := "<region>"

	initAndReturnSingletonInstanceWithAPIKey(authToken, guid, region)

	createEnvironment("environmentId", "environmentName", "desc", "tags", "#FDD13A")
	createCollection("collectionId", "collectionName", "desc", "tags")
	createSegment("segmentName", "segmentId", "desc", "tags", "email", "endsWith", []string{"@in.ibm.com"})
	createFeature("environmentId", "booleanFeatureName", "booleanFeatureId", "desc", "BOOLEAN", "true", "false", "tags", []string{"segmentId"}, 1, "collectionId", "true", "")
	createFeature("environmentId", "numberFeatureName", "numberFeatureId", "desc", "NUMERIC", "1", "2", "tags", []string{"segmentId"}, 1, "collectionId", "3", "")
	createFeature("environmentId", "stringTextFeatureName", "stringTextFeatureId", "desc", "STRING", "enabled", "disabled", "tags", []string{"segmentId"}, 1, "collectionId", "segmentVal", "TEXT")

	featureEnabledValMap := make(map[string]interface{})
	featureEnabledValMap["key"] = "enabled"
	featureDisabledValMap := make(map[string]interface{})
	featureDisabledValMap["key"] = "disabled"
	featureSegmentValMap := make(map[string]interface{})
	featureSegmentValMap["key"] = "segmentVal"

	createFeature("environmentId", "stringJsonFeatureName", "stringJsonFeatureId", "desc", "STRING", featureEnabledValMap, featureDisabledValMap, "tags", []string{"segmentId"}, 1, "collectionId", featureSegmentValMap, "JSON")
	createFeature("environmentId", "stringYamlFeatureName", "stringYamlFeatureId", "desc", "STRING", "---\nkey: enabled", "---\nkey: disabled", "tags", []string{"segmentId"}, 1, "collectionId", "---\nkey: segmentVal\n", "YAML")
	createProperty("environmentId", "booleanPropertyName", "booleanPropertyId", "desc", "BOOLEAN", "true", "tags", []string{"segmentId"}, "collectionId", 2, "true", "")
	createProperty("environmentId", "numberPropertyName", "numberPropertyId", "desc", "NUMERIC", "2", "tags", []string{"segmentId"}, "collectionId", 2, "4", "")
	createProperty("environmentId", "stringTextPropertyName", "stringTextPropertyId", "desc", "STRING", "propertyVal", "tags", []string{"segmentId"}, "collectionId", 2, "segmentVal", "TEXT")

	propertyValMap := make(map[string]interface{})
	propertyValMap["key"] = "enabled"
	propertySegmentValMap := make(map[string]interface{})
	propertySegmentValMap["key"] = "segmentVal"

	createProperty("environmentId", "stringJsonPropertyName", "stringJsonPropertyId", "desc", "STRING", propertyValMap, "tags", []string{"segmentId"}, "collectionId", 2, propertySegmentValMap, "JSON")
	createProperty("environmentId", "stringYamlPropertyName", "stringYamlPropertyId", "desc", "STRING", "---\nkey: propertyVal", "tags", []string{"segmentId"}, "collectionId", 2, "---\nkey: segmentVal", "YAML")

	toggleFeature("environmentId", "booleanFeatureId", true)

	getEnvironments()
	getCollections()
	getFeatures("environmentId")
	getSegments()
	getProperties("environmentId")

	updateFeature("environmentId", "numberFeatureId", "numberFeatureName", "updatedDesc", "1", "1", "tags", []string{}, 1, "collectionId", "2", true)
	updateCollection("collectionId", "collectionName", "updatedDesc", "updatedTags")
	updateSegment("segmentId", "segmentName", "updatedDesc", "updatedTags", "email", "endsWith", []string{"@in.ibm.com"})
	updateProperty("environmentId", "booleanPropertyName", "booleanPropertyId", "updatedDescBoolean", "true", "updatedTags", []string{"segmentId"}, "collectionId", 2, "true")
	updateEnvironment("environmentId", "environmentName", "updatedDesc", "tags", "#FDD13A")

	getEnvironment("environmentId")
	getCollection("collectionId")
	getFeature("environmentId", "booleanFeatureId")
	getProperty("environmentId", "booleanPropertyId")
	getSegment("segmentId")

	patchFeature("environmentId", "booleanFeatureId", "booleanFeatureName", "patchedDesc", "1", "12", "tag", []string{}, 1, "2")
	patchProperty("environmentId", "numberPropertyName", "numberPropertyId", "desc", "1", "tags", []string{"segmentId"}, 2, "2")

	deleteFeature("environmentId", "numberFeatureId")
	deleteFeature("environmentId", "booleanFeatureId")
	deleteFeature("environmentId", "stringTextFeatureId")
	deleteFeature("environmentId", "stringJsonFeatureId")
	deleteFeature("environmentId", "stringYamlFeatureId")
	deleteProperty("environmentId", "numberPropertyId")
	deleteProperty("environmentId", "booleanPropertyId")
	deleteProperty("environmentId", "stringTextPropertyId")
	deleteProperty("environmentId", "stringJsonPropertyId")
	deleteProperty("environmentId", "stringYamlPropertyId")
	deleteSegment("segmentId")
	deleteCollection("collectionId")
	deleteEnvironment("environmentId")
}

// Create examples

func createEnvironment(environmentId string, name string, description string, tags string, colorCode string) {
	fmt.Println("createEnvironment")
	createEnvironmentOptionsModel := appConfigurationServiceInstance.NewCreateEnvironmentOptions(name, environmentId)
	createEnvironmentOptionsModel.SetDescription(description)
	createEnvironmentOptionsModel.SetTags(tags)
	createEnvironmentOptionsModel.SetColorCode(colorCode)
	result, response, error := appConfigurationServiceInstance.CreateEnvironment(createEnvironmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.EnvironmentID)
}

func createCollection(collectionId string, name string, description string, tags string) {
	fmt.Println("createCollection")
	createCollectionOptionsModel := appConfigurationServiceInstance.NewCreateCollectionOptions(name, collectionId)
	createCollectionOptionsModel.SetDescription(description)
	createCollectionOptionsModel.SetTags(tags)
	result, response, error := appConfigurationServiceInstance.CreateCollection(createCollectionOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.CollectionID)
}

func createSegment(name string, id string, description string, tags string, attributeName string, operator string, values []string) {
	fmt.Println("createSegment")
	ruleArray, _ := appConfigurationServiceInstance.NewRule(attributeName, operator, values)
	createSegmentOptionsModel := appConfigurationServiceInstance.NewCreateSegmentOptions()
	createSegmentOptionsModel.SetName(name)
	createSegmentOptionsModel.SetDescription(description)
	createSegmentOptionsModel.SetTags(tags)
	createSegmentOptionsModel.SetSegmentID(id)
	createSegmentOptionsModel.SetRules([]appconfigurationv1.Rule{*ruleArray})
	result, response, error := appConfigurationServiceInstance.CreateSegment(createSegmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.SegmentID)
}

func createFeature(environmentId string, name string, id string, description string, typeOfFeature string, enabledValue interface{}, disabledValue interface{}, tags string, segments []string, order int64, collectionId string, value interface{}, format string) {
	fmt.Println("createFeature")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionRef(collectionId)
	createFeatureOptionsModel := appConfigurationServiceInstance.NewCreateFeatureOptions(environmentId, name, id, typeOfFeature, enabledValue, disabledValue)
	createFeatureOptionsModel.SetTags(tags)
	createFeatureOptionsModel.SetDescription(description)
	createFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	createFeatureOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionArray})
	if len(format) != 0 {
		createFeatureOptionsModel.SetFormat(format)
	}
	result, response, error := appConfigurationServiceInstance.CreateFeature(createFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.FeatureID)
}

func createProperty(environmentId string, name string, propertyId string, description string, typeOfProperty string, valueOfProperty interface{}, tags string, segments []string, collectionId string, order int64, value interface{}, format string) {
	fmt.Println("createProperty")
	ruleArray2, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray2}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionRef(collectionId)
	createPropertyOptionsModel := appConfigurationServiceInstance.NewCreatePropertyOptions(environmentId, name, propertyId, typeOfProperty, valueOfProperty)
	createPropertyOptionsModel.SetTags(tags)
	createPropertyOptionsModel.SetDescription(description)
	createPropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	createPropertyOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionArray})
	if len(format) != 0 {
		createPropertyOptionsModel.SetFormat(format)
	}
	result, response, error := appConfigurationServiceInstance.CreateProperty(createPropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PropertyID)
}

// Update examples

func updateEnvironment(environmentId string, name string, description string, tags string, colorCode string) {
	fmt.Println("updateEnvironment")
	updateEnvironmentOptionsModel := appConfigurationServiceInstance.NewUpdateEnvironmentOptions(environmentId)
	updateEnvironmentOptionsModel.SetName(name)
	updateEnvironmentOptionsModel.SetDescription(description)
	updateEnvironmentOptionsModel.SetTags(tags)
	updateEnvironmentOptionsModel.SetColorCode(colorCode)
	result, response, error := appConfigurationServiceInstance.UpdateEnvironment(updateEnvironmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Description)
}

func updateFeature(environmentId string, id string, name string, description string, enabledValue string, disabledValue string, tags string, segments []string, order int64, collectionId string, value string, deletedFlag bool) {
	fmt.Println("updateFeatureWithNumberValue")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionRef(collectionId)
	updateFeatureOptionsModel := appConfigurationServiceInstance.NewUpdateFeatureOptions(environmentId, id)
	updateFeatureOptionsModel.SetName(name)
	updateFeatureOptionsModel.SetDescription(description)
	updateFeatureOptionsModel.SetTags(tags)
	updateFeatureOptionsModel.SetDisabledValue(disabledValue)
	updateFeatureOptionsModel.SetEnabledValue(enabledValue)
	updateFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	updateFeatureOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionArray})
	result, response, error := appConfigurationServiceInstance.UpdateFeature(updateFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func updateCollection(collectionId string, name string, description string, tags string) {
	fmt.Println("updateCollection")
	updateCollectionOptionsModel := appConfigurationServiceInstance.NewUpdateCollectionOptions(collectionId)
	updateCollectionOptionsModel.SetName(name)
	updateCollectionOptionsModel.SetTags(tags)
	updateCollectionOptionsModel.SetDescription(description)
	result, response, error := appConfigurationServiceInstance.UpdateCollection(updateCollectionOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Description)
}

func patchFeature(environmentId string, id string, name string, description string, enabledValue string, disabledValue string, tags string, segments []string, order int64, value string) {
	fmt.Println("patchFeatureWithNumberValue")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	patchFeatureOptionsModel := appConfigurationServiceInstance.NewUpdateFeatureValuesOptions(environmentId, id)
	patchFeatureOptionsModel.SetName(name)
	patchFeatureOptionsModel.SetDescription(description)
	patchFeatureOptionsModel.SetTags(tags)
	patchFeatureOptionsModel.SetDisabledValue(disabledValue)
	patchFeatureOptionsModel.SetEnabledValue(enabledValue)
	patchFeatureOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	result, response, error := appConfigurationServiceInstance.UpdateFeatureValues(patchFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func updateSegment(segmentId string, name string, description string, tags string, attributeName string, operator string, values []string) {
	fmt.Println("updateSegment")
	ruleArray, _ := appConfigurationServiceInstance.NewRule(attributeName, operator, values)
	updateSegmentOptionsModel := appConfigurationServiceInstance.NewUpdateSegmentOptions(segmentId)
	updateSegmentOptionsModel.SetName(name)
	updateSegmentOptionsModel.SetDescription(description)
	updateSegmentOptionsModel.SetTags(tags)
	updateSegmentOptionsModel.SetRules([]appconfigurationv1.Rule{*ruleArray})
	result, response, error := appConfigurationServiceInstance.UpdateSegment(updateSegmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func updateProperty(environmentId string, name string, propertyId string, description string, valueOfProperty string, tags string, segments []string, collectionId string, order int64, value string) {
	fmt.Println("updateProperty")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	collectionArray, _ := appConfigurationServiceInstance.NewCollectionRef(collectionId)
	updatePropertyOptionsModel := appConfigurationServiceInstance.NewUpdatePropertyOptions(environmentId, propertyId)
	updatePropertyOptionsModel.SetName(name)
	updatePropertyOptionsModel.SetDescription(description)
	updatePropertyOptionsModel.SetTags(tags)
	updatePropertyOptionsModel.SetValue(valueOfProperty)
	updatePropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	updatePropertyOptionsModel.SetCollections([]appconfigurationv1.CollectionRef{*collectionArray})
	result, response, error := appConfigurationServiceInstance.UpdateProperty(updatePropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PropertyID)
}

func patchProperty(environmentId string, name string, propertyId string, description string, valueOfProperty string, tags string, segments []string, order int64, value string) {
	fmt.Println("patchProperty")
	ruleArray, _ := appConfigurationServiceInstance.NewTargetSegments(segments)
	segmentRuleArray, _ := appConfigurationServiceInstance.NewSegmentRule([]appconfigurationv1.TargetSegments{*ruleArray}, value, order)
	patchPropertyOptionsModel := appConfigurationServiceInstance.NewUpdatePropertyValuesOptions(environmentId, propertyId)
	patchPropertyOptionsModel.SetName(name)
	patchPropertyOptionsModel.SetDescription(description)
	patchPropertyOptionsModel.SetTags(tags)
	patchPropertyOptionsModel.SetValue(valueOfProperty)
	patchPropertyOptionsModel.SetSegmentRules([]appconfigurationv1.SegmentRule{*segmentRuleArray})
	result, response, error := appConfigurationServiceInstance.UpdatePropertyValues(patchPropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.PropertyID)
}

// Delete examples

func deleteCollection(collectionId string) {
	fmt.Println("deleteCollection")
	deleteCollectionOptionsModel := appConfigurationServiceInstance.NewDeleteCollectionOptions(collectionId)
	response, error := appConfigurationServiceInstance.DeleteCollection(deleteCollectionOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteFeature(environmentId string, featureId string) {
	fmt.Println("deleteFeature")
	deleteFeatureOptionsModel := appConfigurationServiceInstance.NewDeleteFeatureOptions(environmentId, featureId)
	response, error := appConfigurationServiceInstance.DeleteFeature(deleteFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteSegment(segmentId string) {
	fmt.Println("deleteSegment")
	deleteSegmentOptionsModel := appConfigurationServiceInstance.NewDeleteSegmentOptions(segmentId)
	response, error := appConfigurationServiceInstance.DeleteSegment(deleteSegmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteProperty(environmentId string, propertyId string) {
	fmt.Println("deleteProperty")
	deletePropertyOptionsModel := appConfigurationServiceInstance.NewDeletePropertyOptions(environmentId, propertyId)
	response, error := appConfigurationServiceInstance.DeleteProperty(deletePropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

func deleteEnvironment(environmentId string) {
	fmt.Println("deleteEnvironment")
	deleteEnvironmentOptionsModel := appConfigurationServiceInstance.NewDeleteEnvironmentOptions(environmentId)
	response, error := appConfigurationServiceInstance.DeleteEnvironment(deleteEnvironmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
}

// List examples

func getCollections() {
	fmt.Println("getCollections")
	getCollectionsOptionsModel := appConfigurationServiceInstance.NewListCollectionsOptions()
	getCollectionsOptionsModel.SetExpand(true)
	result, response, error := appConfigurationServiceInstance.ListCollections(getCollectionsOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Collections))
}

func getFeatures(environmentId string) {
	fmt.Println("getFeatures")
	getFeaturesOptionsModel := appConfigurationServiceInstance.NewListFeaturesOptions(environmentId)
	result, response, error := appConfigurationServiceInstance.ListFeatures(getFeaturesOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Features))
}

func getSegments() {
	fmt.Println("getSegments")
	getSegmentsOptionsModel := appConfigurationServiceInstance.NewListSegmentsOptions()
	result, response, error := appConfigurationServiceInstance.ListSegments(getSegmentsOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Segments))
}

func getProperties(environmentId string) {
	fmt.Println("getProperties")
	getPropertiesOptionsModel := appConfigurationServiceInstance.NewListPropertiesOptions(environmentId)
	result, response, error := appConfigurationServiceInstance.ListProperties(getPropertiesOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Properties))
}

func getEnvironments() {
	fmt.Println("getEnvironments")
	getEnvironmentsOptionsModel := appConfigurationServiceInstance.NewListEnvironmentsOptions()
	result, response, error := appConfigurationServiceInstance.ListEnvironments(getEnvironmentsOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(len(result.Environments))
}

// Get examples

func getCollection(collectionId string) {
	fmt.Println("getCollection")
	getCollectionOptionsModel := appConfigurationServiceInstance.NewGetCollectionOptions(collectionId)
	result, response, error := appConfigurationServiceInstance.GetCollection(getCollectionOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func getFeature(environmentId string, featureId string) {
	fmt.Println("getFeature")
	getFeatureOptionsModel := appConfigurationServiceInstance.NewGetFeatureOptions(environmentId, featureId)
	result, response, error := appConfigurationServiceInstance.GetFeature(getFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)        // referenced field, so needs to be de-referenced
	fmt.Println(result.EnabledValue) // non-referenced field, so can be used as it is
}

func getSegment(segmentId string) {
	fmt.Println("getSegment")
	getSegmentOptionsModel := appConfigurationServiceInstance.NewGetSegmentOptions(segmentId)
	result, response, error := appConfigurationServiceInstance.GetSegment(getSegmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func getProperty(environmentId string, propertyId string) {
	fmt.Println("getProperty")
	getPropertyOptionsModel := appConfigurationServiceInstance.NewGetPropertyOptions(environmentId, propertyId)
	result, response, error := appConfigurationServiceInstance.GetProperty(getPropertyOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

func getEnvironment(environmentId string) {
	fmt.Println("getEnvironment")
	getEnvironmentOptionsModel := appConfigurationServiceInstance.NewGetEnvironmentOptions(environmentId)
	result, response, error := appConfigurationServiceInstance.GetEnvironment(getEnvironmentOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}

// Toggle Examples

func toggleFeature(environmentId string, featureId string, enableFlag bool) {
	fmt.Println("toggleFeature")
	toggleFeatureOptionsModel := appConfigurationServiceInstance.NewToggleFeatureOptions(environmentId, featureId)
	toggleFeatureOptionsModel.SetEnabled(enableFlag)
	result, response, error := appConfigurationServiceInstance.ToggleFeature(toggleFeatureOptionsModel)
	if error != nil {
		fmt.Println("Error: " + error.Error())
		return
	}
	fmt.Println(response.StatusCode)
	fmt.Println(*result.Name)
}
