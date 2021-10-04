// Package errors defines all the single use as well as reusable errors within preflight
package errors

import "errors"

var ErrNoChecksEnabled = errors.New("no checks have been enabled")
var ErrRequestedCheckNotFound = errors.New("requested check not found")
var ErrRequestedFormatterNotFound = errors.New("requested formatter is not known")
var ErrFormatterNameNotProvided = errors.New("formatter name is required")
var ErrFormattingResults = errors.New("error formatting results")
var ErrFeatureNotImplemented = errors.New("feature not implemented") // TODO remove this ASAP
var ErrInsufficientPosArguments = errors.New("not enough positional arguments")
var ErrNoResponseFormatSpecified = errors.New("no response format specified")
var ErrGetRemoteContainerFailed = errors.New("failed to pull remote container")
var ErrSaveContainerFailed = errors.New("failed to save container tarball")
var ErrExtractingTarball = errors.New("failed to extract tarball")
var ErrCreateTempDir = errors.New("failed to create temporary directory")
var ErrRunContainerFailed = errors.New("failed to run command in container")
var ErrInitializingLogger = errors.New("failed to initialize logger for config")
var ErrImageInspectFailed = errors.New("failed to inspect image")
var ErrImageScanFailed = errors.New("failed to scan image")
var ErrOperatorSdkScorecardFailed = errors.New("failed to run operator-sdk scorecard")
var ErrOperatorSdkBundleValidateFailed = errors.New("failed to run operator-sdk bundle validate")
var ErrInvalidImageName = errors.New("failed to validate the image name")
var ErrNoSocketFound = errors.New("neither podman or docker socket found")
var ErrAlreadyInUnshare = errors.New("unshare is already active")
var ErrDeterminingRelatedImageSchemaVers = errors.New("failed to determine related image schema version")
var ErrNoKubeconfig = errors.New("no environment variable KUBECONFIG could be found")
var ErrUnableToRetrieveUser = errors.New("could not retrieve user info")
var ErrK8sAPICallFailed = errors.New("unable to fetch the requested resource from k8s API server")
var ErrNoValueFoundInViper = errors.New("unable to fetch the requested key from viper")
var ErrEmptyAnnotationFile = errors.New("the annotations file was empty")
var ErrLicensesNotADir = errors.New("licenses is not a directory")
var ErrSupportCmdPromptFailed = errors.New("prompt failed, please try re-running support command")
var ErrEmptyProjectID = errors.New("please enter a non empty project id")
var ErrRemovePFromProjectID = errors.New("please remove leading character p from project id")
var ErrRemoveOSPIDFromProjectID = errors.New("please remove leading character ospid- from project id")
var ErrRemoveSpecialCharFromProjectID = errors.New("please remove all special characters from project id")
var ErrPullRequestURL = errors.New("please enter a valid url: including scheme, host, and path to pull request")
var ErrIndexImageUndefined = errors.New("no environment variable PFLT_INDEXIMAGE could be found")
var ErrNoDockerconfig = errors.New("no environment variable DOCKERCONFIG could be found")
