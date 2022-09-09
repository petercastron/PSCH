Pod::Spec.new do |spec|
  spec.name         = 'psch'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/petercastron/PSCH'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS psch Client'
  spec.source       = { :git => 'https://github.com/petercastron/PSCH.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/psch.framework'

	spec.prepare_command = <<-CMD
    curl https://pschstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/psch.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
