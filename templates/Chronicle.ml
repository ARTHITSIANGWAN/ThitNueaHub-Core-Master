// TNH-MEMORY-TEMPLATE
// Recorder: Chronicle (The Archiver)

stream Chronicle {
    log_level: "VERBOSE"
    storage: "LOCAL_SATCHA_DB"
    
    event Sync_Update {
        timestamp: "AUTO"
        status: "VERIFIED"
        version: "LATEST"
    }

    event Job_Processed {
        job_id: "FROM_MAIN_GO"
        outcome: "SUCCESS"
    }
}

// บันทึกทุกความเคลื่อนไหวลงในฐานข้อมูลจำลอง
