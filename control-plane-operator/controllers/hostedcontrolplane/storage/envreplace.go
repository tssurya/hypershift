package storage

import (
	v1 "k8s.io/api/core/v1"
)

var (
	// map env. variable in CSO Deployment -> name of the image in payload.
	operatorImageRefs = map[string]string{
		"AWS_EBS_DRIVER_OPERATOR_IMAGE":                   "aws-ebs-csi-driver-operator",
		"AWS_EBS_DRIVER_IMAGE":                            "aws-ebs-csi-driver",
		"GCP_PD_DRIVER_OPERATOR_IMAGE":                    "gcp-pd-csi-driver-operator",
		"GCP_PD_DRIVER_IMAGE":                             "gcp-pd-csi-driver",
		"OPENSTACK_CINDER_DRIVER_OPERATOR_IMAGE":          "openstack-cinder-csi-driver-operator",
		"OPENSTACK_CINDER_DRIVER_IMAGE":                   "openstack-cinder-csi-driver",
		"OVIRT_DRIVER_OPERATOR_IMAGE":                     "ovirt-csi-driver-operator",
		"OVIRT_DRIVER_IMAGE":                              "ovirt-csi-driver",
		"MANILA_DRIVER_OPERATOR_IMAGE":                    "csi-driver-manila-operator",
		"MANILA_DRIVER_IMAGE":                             "csi-driver-manila",
		"MANILA_NFS_DRIVER_IMAGE":                         "csi-driver-nfs",
		"PROVISIONER_IMAGE":                               "csi-external-provisioner",
		"ATTACHER_IMAGE":                                  "csi-external-attacher",
		"RESIZER_IMAGE":                                   "csi-external-resizer",
		"SNAPSHOTTER_IMAGE":                               "csi-external-snapshotter",
		"NODE_DRIVER_REGISTRAR_IMAGE":                     "csi-node-driver-registrar",
		"LIVENESS_PROBE_IMAGE":                            "csi-livenessprobe",
		"VSPHERE_PROBLEM_DETECTOR_OPERATOR_IMAGE":         "vsphere-problem-detector",
		"AZURE_DISK_DRIVER_OPERATOR_IMAGE":                "azure-disk-csi-driver-operator",
		"AZURE_DISK_DRIVER_IMAGE":                         "azure-disk-csi-driver",
		"AZURE_FILE_DRIVER_OPERATOR_IMAGE":                "azure-file-csi-driver-operator",
		"AZURE_FILE_DRIVER_IMAGE":                         "azure-file-csi-driver",
		"KUBE_RBAC_PROXY_IMAGE":                           "kube-rbac-proxy",
		"VMWARE_VSPHERE_DRIVER_OPERATOR_IMAGE":            "vsphere-csi-driver-operator",
		"VMWARE_VSPHERE_DRIVER_IMAGE":                     "vsphere-csi-driver",
		"VMWARE_VSPHERE_SYNCER_IMAGE":                     "vsphere-csi-driver-syncer",
		"CLUSTER_CLOUD_CONTROLLER_MANAGER_OPERATOR_IMAGE": "cluster-cloud-controller-manager-operator",
		"SHARED_RESOURCE_DRIVER_OPERATOR_IMAGE":           "csi-driver-shared-resource-operator",
		"SHARED_RESOURCE_DRIVER_IMAGE":                    "csi-driver-shared-resource",
		"SHARED_RESOURCE_DRIVER_WEBHOOK_IMAGE":            "csi-driver-shared-resource-webhook",
		"ALIBABA_DISK_DRIVER_OPERATOR_IMAGE":              "alibaba-disk-csi-driver-operator",
		"ALIBABA_CLOUD_DRIVER_IMAGE":                      "alibaba-cloud-csi-driver",
		"IBM_VPC_BLOCK_DRIVER_OPERATOR_IMAGE":             "ibm-vpc-block-csi-driver-operator",
		"IBM_VPC_BLOCK_DRIVER_IMAGE":                      "ibm-vpc-block-csi-driver",
		"IBM_VPC_NODE_LABEL_UPDATER_IMAGE":                "ibm-vpc-node-label-updater",
		"POWERVS_BLOCK_CSI_DRIVER_OPERATOR_IMAGE":         "powervs-block-csi-driver-operator",
		"POWERVS_BLOCK_CSI_DRIVER_IMAGE":                  "powervs-block-csi-driver",
		"HYPERSHIFT_IMAGE":                                "token-minter",
	}
)

type environmentReplacer struct {
	// map env variable name -> new env. variable value
	values map[string]string
}

func newEnvironmentReplacer() *environmentReplacer {
	return &environmentReplacer{values: map[string]string{}}
}

func (er *environmentReplacer) setOperatorImageReferences(images map[string]string, userImages map[string]string) {
	// `operatorImageRefs` is map from env. var name -> payload image name
	// `images` is map from payload image name -> image URL
	// Create map from env. var name -> image URL
	for envVar, payloadName := range operatorImageRefs {
		// TODO: *_DRIVER_IMAGE and LIVENESS_PROBE_IMAGE for csi-driver-node should be separated
		// from csi-driver-controller and looked up from the userImages
		// strings.HasSuffix(envVar, "_DRIVER_IMAGE") || envVar == "LIVENESS_PROBE_IMAGE"
		if envVar == "NODE_DRIVER_REGISTRAR_IMAGE" {
			if imageURL, ok := userImages[payloadName]; ok {
				er.values[envVar] = imageURL
			}
		} else {
			if imageURL, ok := images[payloadName]; ok {
				er.values[envVar] = imageURL
			}
		}
	}
}

func (er *environmentReplacer) setVersions(ver string) {
	er.values["OPERATOR_IMAGE_VERSION"] = ver
	er.values["OPERAND_IMAGE_VERSION"] = ver
}

func (er *environmentReplacer) replaceEnvVars(envVars []v1.EnvVar) {
	for i := range envVars {
		if value, ok := er.values[envVars[i].Name]; ok {
			envVars[i].Value = value
		}
	}
}
