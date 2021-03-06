package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// AnalyticsProfile analytics profile
// swagger:model AnalyticsProfile
type AnalyticsProfile struct {

	// If a client receives an HTTP response in less than the Satisfactory Latency Threshold, the request is considered Satisfied. It is considered Tolerated if it is not Satisfied and less than Tolerated Latency Factor multiplied by the Satisfactory Latency Threshold. Greater than this number and the client's request is considered Frustrated. Allowed values are 1-30000. Units(MILLISECONDS).
	ApdexResponseThreshold int32 `json:"apdex_response_threshold,omitempty"`

	// Client tolerated response latency factor. Client must receive a response within this factor times the satisfactory threshold (apdex_response_threshold) to be considered tolerated. Allowed values are 1-1000.
	ApdexResponseToleratedFactor *float64 `json:"apdex_response_tolerated_factor,omitempty"`

	// Satisfactory client to Avi Round Trip Time(RTT). Allowed values are 1-2000. Units(MILLISECONDS).
	ApdexRttThreshold int32 `json:"apdex_rtt_threshold,omitempty"`

	// Tolerated client to Avi Round Trip Time(RTT) factor.  It is a multiple of apdex_rtt_tolerated_factor. Allowed values are 1-1000.
	ApdexRttToleratedFactor *float64 `json:"apdex_rtt_tolerated_factor,omitempty"`

	// If a client is able to load a page in less than the Satisfactory Latency Threshold, the PageLoad is considered Satisfied.  It is considered tolerated if it is greater than Satisfied but less than the Tolerated Latency multiplied by Satisifed Latency. Greater than this number and the client's request is considered Frustrated.  A PageLoad includes the time for DNS lookup, download of all HTTP objects, and page render time. Allowed values are 1-30000. Units(MILLISECONDS).
	ApdexRumThreshold int32 `json:"apdex_rum_threshold,omitempty"`

	// Virtual service threshold factor for tolerated Page Load Time (PLT) as multiple of apdex_rum_threshold. Allowed values are 1-1000.
	ApdexRumToleratedFactor *float64 `json:"apdex_rum_tolerated_factor,omitempty"`

	// A server HTTP response is considered Satisfied if latency is less than the Satisfactory Latency Threshold. The response is considered tolerated when it is greater than Satisfied but less than the Tolerated Latency Factor * S_Latency.  Greater than this number and the server response is considered Frustrated. Allowed values are 1-30000. Units(MILLISECONDS).
	ApdexServerResponseThreshold int32 `json:"apdex_server_response_threshold,omitempty"`

	// Server tolerated response latency factor. Servermust response within this factor times the satisfactory threshold (apdex_server_response_threshold) to be considered tolerated. Allowed values are 1-1000.
	ApdexServerResponseToleratedFactor *float64 `json:"apdex_server_response_tolerated_factor,omitempty"`

	// Satisfactory client to Avi Round Trip Time(RTT). Allowed values are 1-2000. Units(MILLISECONDS).
	ApdexServerRttThreshold int32 `json:"apdex_server_rtt_threshold,omitempty"`

	// Tolerated client to Avi Round Trip Time(RTT) factor.  It is a multiple of apdex_rtt_tolerated_factor. Allowed values are 1-1000.
	ApdexServerRttToleratedFactor *float64 `json:"apdex_server_rtt_tolerated_factor,omitempty"`

	// Configure which logs are sent to the Avi Controller from SEs and how they are processed.
	ClientLogConfig *ClientLogConfiguration `json:"client_log_config,omitempty"`

	// Configure to stream logs to an external server. Field introduced in 17.1.1.
	ClientLogStreamingConfig *ClientLogStreamingConfig `json:"client_log_streaming_config,omitempty"`

	// A connection between client and Avi is considered lossy when more than this percentage of out of order packets are received. Allowed values are 1-100. Units(PERCENT).
	ConnLossyOooThreshold int32 `json:"conn_lossy_ooo_threshold,omitempty"`

	// A connection between client and Avi is considered lossy when more than this percentage of packets are retransmitted due to timeout. Allowed values are 1-100. Units(PERCENT).
	ConnLossyTimeoRexmtThreshold int32 `json:"conn_lossy_timeo_rexmt_threshold,omitempty"`

	// A connection between client and Avi is considered lossy when more than this percentage of packets are retransmitted. Allowed values are 1-100. Units(PERCENT).
	ConnLossyTotalRexmtThreshold int32 `json:"conn_lossy_total_rexmt_threshold,omitempty"`

	// A client connection is considered lossy when percentage of times a packet could not be trasmitted due to TCP zero window is above this threshold. Allowed values are 0-100. Units(PERCENT).
	ConnLossyZeroWinSizeEventThreshold int32 `json:"conn_lossy_zero_win_size_event_threshold,omitempty"`

	// A connection between Avi and server is considered lossy when more than this percentage of out of order packets are received. Allowed values are 1-100. Units(PERCENT).
	ConnServerLossyOooThreshold int32 `json:"conn_server_lossy_ooo_threshold,omitempty"`

	// A connection between Avi and server is considered lossy when more than this percentage of packets are retransmitted due to timeout. Allowed values are 1-100. Units(PERCENT).
	ConnServerLossyTimeoRexmtThreshold int32 `json:"conn_server_lossy_timeo_rexmt_threshold,omitempty"`

	// A connection between Avi and server is considered lossy when more than this percentage of packets are retransmitted. Allowed values are 1-100. Units(PERCENT).
	ConnServerLossyTotalRexmtThreshold int32 `json:"conn_server_lossy_total_rexmt_threshold,omitempty"`

	// A server connection is considered lossy when percentage of times a packet could not be trasmitted due to TCP zero window is above this threshold. Allowed values are 0-100. Units(PERCENT).
	ConnServerLossyZeroWinSizeEventThreshold int32 `json:"conn_server_lossy_zero_win_size_event_threshold,omitempty"`

	// User defined description for the object.
	Description string `json:"description,omitempty"`

	// Disable node (service engine) level analytics forvs metrics.
	DisableSeAnalytics bool `json:"disable_se_analytics,omitempty"`

	// Disable analytics on backend servers. This may be desired in container environment when there are large number of  ephemeral servers.
	DisableServerAnalytics bool `json:"disable_server_analytics,omitempty"`

	// Exclude client closed connection before an HTTP request could be completed from being classified as an error.
	ExcludeClientCloseBeforeRequestAsError bool `json:"exclude_client_close_before_request_as_error,omitempty"`

	// Exclude dns policy drops from the list of errors. Field introduced in 17.2.2.
	ExcludeDNSPolicyDropAsSignificant bool `json:"exclude_dns_policy_drop_as_significant,omitempty"`

	// Exclude queries to GSLB services that are operationally down from the list of errors.
	ExcludeGsDownAsError bool `json:"exclude_gs_down_as_error,omitempty"`

	// List of HTTP status codes to be excluded from being classified as an error.  Error connections or responses impacts health score, are included as significant logs, and may be classified as part of a DoS attack.
	ExcludeHTTPErrorCodes []int64 `json:"exclude_http_error_codes,omitempty,omitempty"`

	// Exclude dns queries to domains outside the domains configured in the DNS application profile from the list of errors.
	ExcludeInvalidDNSDomainAsError bool `json:"exclude_invalid_dns_domain_as_error,omitempty"`

	// Exclude invalid dns queries from the list of errors.
	ExcludeInvalidDNSQueryAsError bool `json:"exclude_invalid_dns_query_as_error,omitempty"`

	// Exclude queries to domains that did not have configured services/records from the list of errors.
	ExcludeNoDNSRecordAsError bool `json:"exclude_no_dns_record_as_error,omitempty"`

	// Exclude queries to GSLB services that have no available members from the list of errors.
	ExcludeNoValidGsMemberAsError bool `json:"exclude_no_valid_gs_member_as_error,omitempty"`

	// Exclude persistence server changed while load balancing' from the list of errors.
	ExcludePersistenceChangeAsError bool `json:"exclude_persistence_change_as_error,omitempty"`

	// Exclude server dns error response from the list of errors.
	ExcludeServerDNSErrorAsError bool `json:"exclude_server_dns_error_as_error,omitempty"`

	// Exclude server TCP reset from errors.  It is common for applications like MS Exchange.
	ExcludeServerTCPResetAsError bool `json:"exclude_server_tcp_reset_as_error,omitempty"`

	// Exclude 'server unanswered syns' from the list of errors.
	ExcludeSynRetransmitAsError bool `json:"exclude_syn_retransmit_as_error,omitempty"`

	// Exclude TCP resets by client from the list of potential errors.
	ExcludeTCPResetAsError bool `json:"exclude_tcp_reset_as_error,omitempty"`

	// Exclude unsupported dns queries from the list of errors.
	ExcludeUnsupportedDNSQueryAsError bool `json:"exclude_unsupported_dns_query_as_error,omitempty"`

	// Time window (in secs) within which only unique health change events should occur.
	HsEventThrottleWindow int32 `json:"hs_event_throttle_window,omitempty"`

	// Maximum penalty that may be deducted from health score for anomalies. Allowed values are 0-100.
	HsMaxAnomalyPenalty int32 `json:"hs_max_anomaly_penalty,omitempty"`

	// Maximum penalty that may be deducted from health score for high resource utilization. Allowed values are 0-100.
	HsMaxResourcesPenalty int32 `json:"hs_max_resources_penalty,omitempty"`

	// Maximum penalty that may be deducted from health score based on security assessment. Allowed values are 0-100.
	HsMaxSecurityPenalty int32 `json:"hs_max_security_penalty,omitempty"`

	// DoS connection rate below which the DoS security assessment will not kick in.
	HsMinDosRate int32 `json:"hs_min_dos_rate,omitempty"`

	// Adds free performance score credits to health score. It can be used for compensating health score for known slow applications. Allowed values are 0-100.
	HsPerformanceBoost int32 `json:"hs_performance_boost,omitempty"`

	// Threshold number of connections in 5min, below which apdexr, apdexc, rum_apdex, and other network quality metrics are not computed.
	HsPscoreTrafficThresholdL4Client *float64 `json:"hs_pscore_traffic_threshold_l4_client,omitempty"`

	// Threshold number of connections in 5min, below which apdexr, apdexc, rum_apdex, and other network quality metrics are not computed.
	HsPscoreTrafficThresholdL4Server *float64 `json:"hs_pscore_traffic_threshold_l4_server,omitempty"`

	// Score assigned when the certificate has expired. Allowed values are 0-5.
	HsSecurityCertscoreExpired float64 `json:"hs_security_certscore_expired,omitempty"`

	// Score assigned when the certificate expires in more than 30 days. Allowed values are 0-5.
	HsSecurityCertscoreGt30d *float64 `json:"hs_security_certscore_gt30d,omitempty"`

	// Score assigned when the certificate expires in less than or equal to 7 days. Allowed values are 0-5.
	HsSecurityCertscoreLe07d *float64 `json:"hs_security_certscore_le07d,omitempty"`

	// Score assigned when the certificate expires in less than or equal to 30 days. Allowed values are 0-5.
	HsSecurityCertscoreLe30d *float64 `json:"hs_security_certscore_le30d,omitempty"`

	// Penalty for allowing certificates with invalid chain. Allowed values are 0-5.
	HsSecurityChainInvalidityPenalty *float64 `json:"hs_security_chain_invalidity_penalty,omitempty"`

	// Score assigned when the minimum cipher strength is 0 bits. Allowed values are 0-5.
	HsSecurityCipherscoreEq000b float64 `json:"hs_security_cipherscore_eq000b,omitempty"`

	// Score assigned when the minimum cipher strength is greater than equal to 128 bits. Allowed values are 0-5.
	HsSecurityCipherscoreGe128b *float64 `json:"hs_security_cipherscore_ge128b,omitempty"`

	// Score assigned when the minimum cipher strength is less than 128 bits. Allowed values are 0-5.
	HsSecurityCipherscoreLt128b *float64 `json:"hs_security_cipherscore_lt128b,omitempty"`

	// Score assigned when no algorithm is used for encryption. Allowed values are 0-5.
	HsSecurityEncalgoScoreNone float64 `json:"hs_security_encalgo_score_none,omitempty"`

	// Score assigned when RC4 algorithm is used for encryption. Allowed values are 0-5.
	HsSecurityEncalgoScoreRc4 *float64 `json:"hs_security_encalgo_score_rc4,omitempty"`

	// Penalty for not enabling HSTS. Allowed values are 0-5.
	HsSecurityHstsPenalty *float64 `json:"hs_security_hsts_penalty,omitempty"`

	// Penalty for allowing non-PFS handshakes. Allowed values are 0-5.
	HsSecurityNonpfsPenalty *float64 `json:"hs_security_nonpfs_penalty,omitempty"`

	// Deprecated. Allowed values are 0-5.
	HsSecuritySelfsignedcertPenalty *float64 `json:"hs_security_selfsignedcert_penalty,omitempty"`

	// Score assigned when supporting SSL3.0 encryption protocol. Allowed values are 0-5.
	HsSecuritySsl30Score *float64 `json:"hs_security_ssl30_score,omitempty"`

	// Score assigned when supporting TLS1.0 encryption protocol. Allowed values are 0-5.
	HsSecurityTLS10Score *float64 `json:"hs_security_tls10_score,omitempty"`

	// Score assigned when supporting TLS1.1 encryption protocol. Allowed values are 0-5.
	HsSecurityTLS11Score *float64 `json:"hs_security_tls11_score,omitempty"`

	// Score assigned when supporting TLS1.2 encryption protocol. Allowed values are 0-5.
	HsSecurityTLS12Score *float64 `json:"hs_security_tls12_score,omitempty"`

	// Penalty for allowing weak signature algorithm(s). Allowed values are 0-5.
	HsSecurityWeakSignatureAlgoPenalty *float64 `json:"hs_security_weak_signature_algo_penalty,omitempty"`

	// The name of the analytics profile.
	// Required: true
	Name string `json:"name"`

	// List of HTTP status code ranges to be excluded from being classified as an error.
	Ranges []*HttpstatusRange `json:"ranges,omitempty"`

	// Block of HTTP response codes to be excluded from being classified as an error. Enum options - AP_HTTP_RSP_4XX, AP_HTTP_RSP_5XX.
	RespCodeBlock []string `json:"resp_code_block,omitempty"`

	//  It is a reference to an object of type Tenant.
	TenantRef string `json:"tenant_ref,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// UUID of the analytics profile.
	UUID string `json:"uuid,omitempty"`
}
