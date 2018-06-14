package wrapper

import (
	"testing"
	"strings"
	"log"
	"runtime"
	"path/filepath"
)

var test_output = `<Iteration>
    <Iteration_iter-num>2</Iteration_iter-num>
    <Iteration_query-ID>Query_2</Iteration_query-ID>
    <Iteration_query-def>gi|16080617|ref|NP_391444.1| membrane bound lipoprotein [Bacillus subtilis subsp. subtilis str. 168]</Iteration_query-def>
    <Iteration_query-len>102</Iteration_query-len>
    <Iteration_hits>
        <Hit>
            <Hit_num>1</Hit_num>
            <Hit_id>gnl|BL_ORD_ID|1</Hit_id>
            <Hit_def>gi|308175296|ref|YP_003922001.1| membrane bound lipoprotein [Bacillus amyloliquefaciens DSM 7]</Hit_def>
            <Hit_accession>1</Hit_accession>
            <Hit_len>100</Hit_len>
            <Hit_hsps>
                <Hsp>
                    <Hsp_num>1</Hsp_num>
                    <Hsp_bit-score>139.428</Hsp_bit-score>
                    <Hsp_score>350</Hsp_score>
                    <Hsp_evalue>1.99275e-46</Hsp_evalue>
                    <Hsp_query-from>1</Hsp_query-from>
                    <Hsp_query-to>102</Hsp_query-to>
                    <Hsp_hit-from>1</Hsp_hit-from>
                    <Hsp_hit-to>100</Hsp_hit-to>
                    <Hsp_query-frame>0</Hsp_query-frame>
                    <Hsp_hit-frame>0</Hsp_hit-frame>
                    <Hsp_identity>69</Hsp_identity>
                    <Hsp_positive>81</Hsp_positive>
                    <Hsp_gaps>2</Hsp_gaps>
                    <Hsp_align-len>102</Hsp_align-len>
                    <Hsp_qseq>MKKFIALLFFILLLSGCGVNSQKSQGEDVSPDSNIETKEGTYVGLADTHTIEVTVDNEPVSLDITEESTSDLDKFNSGDKVTITYEKNDEGQLLLKDIERAN</Hsp_qseq>
                    <Hsp_hseq>MKKIFGCLFFILLLAGCGVTNEKSQGEDAG--EKLVTKEGTYVGLADTHTIEVTVDHEPVSFDITEESADDVKNLNNGEKVTVKYQKNSKGQLVLKDIEPAN</Hsp_hseq>
                    <Hsp_midline>MKK    LFFILLL+GCGV ++KSQGED      + TKEGTYVGLADTHTIEVTVD+EPVS DITEES  D+   N+G+KVT+ Y+KN +GQL+LKDIE AN</Hsp_midline>
                </Hsp>
            </Hit_hsps>
        </Hit>
        <Hit>
            <Hit_num>2</Hit_num>
            <Hit_id>gnl|BL_ORD_ID|2</Hit_id>
            <Hit_def>gi|375363999|ref|YP_005132038.1| lytA gene product [Bacillus amyloliquefaciens subsp. plantarum CAU B946]</Hit_def>
            <Hit_accession>2</Hit_accession>
            <Hit_len>105</Hit_len>
            <Hit_hsps>
                <Hsp>
                    <Hsp_num>1</Hsp_num>
                    <Hsp_bit-score>88.9669</Hsp_bit-score>
                    <Hsp_score>219</Hsp_score>
                    <Hsp_evalue>6.94052e-27</Hsp_evalue>
                    <Hsp_query-from>1</Hsp_query-from>
                    <Hsp_query-to>101</Hsp_query-to>
                    <Hsp_hit-from>1</Hsp_hit-from>
                    <Hsp_hit-to>104</Hsp_hit-to>
                    <Hsp_query-frame>0</Hsp_query-frame>
                    <Hsp_hit-frame>0</Hsp_hit-frame>
                    <Hsp_identity>48</Hsp_identity>
                    <Hsp_positive>69</Hsp_positive>
                    <Hsp_gaps>5</Hsp_gaps>
                    <Hsp_align-len>105</Hsp_align-len>
                    <Hsp_qseq>MKKFIALLFFILL----LSGCGVNSQKSQGEDVSPDSNIETKEGTYVGLADTHTIEVTVDNEPVSLDITEESTSDLDKFNSGDKVTITYEKNDEGQLLLKDIERA</Hsp_qseq>
                    <Hsp_hseq>MKKTIAASFLILLFSVVLAACGTAEQSKKGSG-SSENQAQKETAYYVGMADTHTIEVKVDDQPVSFEFSDDFSDVLNKFSENDKVSITYFTNDKGQKEIKEIEKA</Hsp_hseq>
                    <Hsp_midline>MKK IA  F ILL    L+ CG   Q  +G   S ++  + +   YVG+ADTHTIEV VD++PVS + +++ +  L+KF+  DKV+ITY  ND+GQ  +K+IE+A</Hsp_midline>
                </Hsp>
            </Hit_hsps>
        </Hit>
        <Hit>
            <Hit_num>3</Hit_num>
            <Hit_id>gnl|BL_ORD_ID|3</Hit_id>
            <Hit_def>gi|154687679|ref|YP_001422840.1| LytA [Bacillus amyloliquefaciens FZB42]</Hit_def>
            <Hit_accession>3</Hit_accession>
            <Hit_len>105</Hit_len>
            <Hit_hsps>
                <Hsp>
                    <Hsp_num>1</Hsp_num>
                    <Hsp_bit-score>88.9669</Hsp_bit-score>
                    <Hsp_score>219</Hsp_score>
                    <Hsp_evalue>8.41012e-27</Hsp_evalue>
                    <Hsp_query-from>1</Hsp_query-from>
                    <Hsp_query-to>101</Hsp_query-to>
                    <Hsp_hit-from>1</Hsp_hit-from>
                    <Hsp_hit-to>104</Hsp_hit-to>
                    <Hsp_query-frame>0</Hsp_query-frame>
                    <Hsp_hit-frame>0</Hsp_hit-frame>
                    <Hsp_identity>48</Hsp_identity>
                    <Hsp_positive>69</Hsp_positive>
                    <Hsp_gaps>5</Hsp_gaps>
                    <Hsp_align-len>105</Hsp_align-len>
                    <Hsp_qseq>MKKFIALLFFILL----LSGCGVNSQKSQGEDVSPDSNIETKEGTYVGLADTHTIEVTVDNEPVSLDITEESTSDLDKFNSGDKVTITYEKNDEGQLLLKDIERA</Hsp_qseq>
                    <Hsp_hseq>MKKTIAASFLILLFSVVLAACGTADQSKKGSG-SSENQAQKETAYYVGMADTHTIEVKVDDQPVSFEFSDDFSDVLNKFSENDKVSITYFTNDKGQKEIKEIEKA</Hsp_hseq>
                    <Hsp_midline>MKK IA  F ILL    L+ CG   Q  +G   S ++  + +   YVG+ADTHTIEV VD++PVS + +++ +  L+KF+  DKV+ITY  ND+GQ  +K+IE+A</Hsp_midline>
                </Hsp>
            </Hit_hsps>
        </Hit>
        <Hit>
            <Hit_num>4</Hit_num>
            <Hit_id>gnl|BL_ORD_ID|4</Hit_id>
            <Hit_def>gi|311070071|ref|YP_003974994.1| unnamed protein product [Bacillus atrophaeus 1942]</Hit_def>
            <Hit_accession>4</Hit_accession>
            <Hit_len>105</Hit_len>
            <Hit_hsps>
                <Hsp>
                    <Hsp_num>1</Hsp_num>
                    <Hsp_bit-score>83.1889</Hsp_bit-score>
                    <Hsp_score>204</Hsp_score>
                    <Hsp_evalue>1.37847e-24</Hsp_evalue>
                    <Hsp_query-from>1</Hsp_query-from>
                    <Hsp_query-to>100</Hsp_query-to>
                    <Hsp_hit-from>1</Hsp_hit-from>
                    <Hsp_hit-to>103</Hsp_hit-to>
                    <Hsp_query-frame>0</Hsp_query-frame>
                    <Hsp_hit-frame>0</Hsp_hit-frame>
                    <Hsp_identity>45</Hsp_identity>
                    <Hsp_positive>66</Hsp_positive>
                    <Hsp_gaps>5</Hsp_gaps>
                    <Hsp_align-len>104</Hsp_align-len>
                    <Hsp_qseq>MKKFIALLFFILL----LSGCGVNSQKSQGEDVSPDSNIETKEGTYVGLADTHTIEVTVDNEPVSLDITEESTSDLDKFNSGDKVTITYEKNDEGQLLLKDIER</Hsp_qseq>
                    <Hsp_hseq>MKKNVASSFLILLFSIILAACGTAEQSKEG-NGSSSSQVQNETAYYVGMADTHTIEVKIDDQPVSFEFTDDFSEILNEFEENDKVNISYLTNDKGQKELTEIEK</Hsp_hseq>
                    <Hsp_midline>MKK +A  F ILL    L+ CG   Q  +G + S  S ++ +   YVG+ADTHTIEV +D++PVS + T++ +  L++F   DKV I+Y  ND+GQ  L +IE+</Hsp_midline>
                </Hsp>
            </Hit_hsps>
        </Hit>
        <Hit>
            <Hit_num>5</Hit_num>
            <Hit_id>gnl|BL_ORD_ID|15</Hit_id>
            <Hit_def>gi|332258565|ref|XP_003278367.1| PREDICTED: UPF0764 protein C16orf89-like [Nomascus leucogenys]</Hit_def>
            <Hit_accession>15</Hit_accession>
            <Hit_len>132</Hit_len>
            <Hit_hsps>
                <Hsp>
                    <Hsp_num>1</Hsp_num>
                    <Hsp_bit-score>15.779</Hsp_bit-score>
                    <Hsp_score>29</Hsp_score>
                    <Hsp_evalue>7.12269</Hsp_evalue>
                    <Hsp_query-from>60</Hsp_query-from>
                    <Hsp_query-to>84</Hsp_query-to>
                    <Hsp_hit-from>80</Hsp_hit-from>
                    <Hsp_hit-to>104</Hsp_hit-to>
                    <Hsp_query-frame>0</Hsp_query-frame>
                    <Hsp_hit-frame>0</Hsp_hit-frame>
                    <Hsp_identity>7</Hsp_identity>
                    <Hsp_positive>11</Hsp_positive>
                    <Hsp_gaps>0</Hsp_gaps>
                    <Hsp_align-len>25</Hsp_align-len>
                    <Hsp_qseq>VSLDITEESTSDLDKFNSGDKVTIT</Hsp_qseq>
                    <Hsp_hseq>VEMGFLHVGQAGLELVTSGDPPTLT</Hsp_hseq>
                    <Hsp_midline>V +       + L+   SGD  T+T</Hsp_midline>
                </Hsp>
            </Hit_hsps>
        </Hit>
    </Iteration_hits>
    <Iteration_stat>
        <Statistics>
            <Statistics_db-num>20</Statistics_db-num>
            <Statistics_db-len>6406</Statistics_db-len>
            <Statistics_hsp-len>38</Statistics_hsp-len>
            <Statistics_eff-space>361344</Statistics_eff-space>
            <Statistics_kappa>0.041</Statistics_kappa>
            <Statistics_lambda>0.267</Statistics_lambda>
            <Statistics_entropy>0.14</Statistics_entropy>
        </Statistics>
    </Iteration_stat>
</Iteration>`

func TestParseFromXMLReader(t *testing.T) {
	r := strings.NewReader(test_output)
	c:= make(chan TBlastXHit)
	go ParseFromXMLReader(r, c)
	for h:= range c {
		log.Println(h)
	}
}

func TestReadTBlastXOutput(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	p, _ := filepath.Split(filename)
	r := filepath.Join(p, "tblastx_output_test.xml")
	ParseHitTBlastXOutputXML(r)
}