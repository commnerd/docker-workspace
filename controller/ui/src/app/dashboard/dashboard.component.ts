import { Component, OnInit } from '@angular/core'
import { FormControl } from '@angular/forms'
import { Proxy } from '../models/proxy'

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  proxyControls: Array<Object>

  constructor() {}

  ngOnInit() {
  	this.proxyControls = []
  	this.addProxy(new Proxy)
  }

  addProxy(proxy: Proxy) {
  	let control = {
  		port: new FormControl(proxy.port),
  		env_base_path: new FormControl(proxy.env_base_path),
  		site_base_path: new FormControl(proxy.site_base_path)
  	}
  	this.proxyControls.push(control)
  }

}
