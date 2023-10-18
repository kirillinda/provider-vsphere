package vSphereVappContainer

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vsphere_vapp_container", func(r *config.Resource) {
		r.ShortGroup = "Virtual_Machine"
		r.Kind = "VSphereVappContainer"
		r.Version = "v1alpha1"
	})
}