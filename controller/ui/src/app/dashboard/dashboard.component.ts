import { Component, OnInit } from '@angular/core'
import { FormBuilder } from '@angular/forms'

import { Proxy } from '../models/proxy'

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {

  dashboardForm = this._fb.group({
    proxies: this.fb.array({
      port: new FormControl(proxy.port),
      env_base_path: new FormControl(proxy.env_base_path),
      site_base_path: new FormControl(proxy.site_base_path)
    })
  })

  constructor(private _fb: FormBuilder) {}

  ngOnInit() {
  	this.proxyControls = []
  	this.addProxy(new Proxy)
  }

  addProxy(proxy: Proxy) {
  	let control = 
  	this.proxyControls.push(control)
  }

}
